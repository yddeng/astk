package astks

import (
	"fmt"
	"github.com/yddeng/astk/pkg/common"
	"github.com/yddeng/astk/pkg/protocol"
	"github.com/yddeng/astk/pkg/types"
	"github.com/yddeng/dnet/drpc"
	"log"
	"strings"
	"syscall"
)

type ProcessConfig struct {
	Name    string `json:"name"`
	Context string `json:"context"`
}

type ProcessState struct {
	// 状态
	Status types.ProcStatus `json:"status"`
	// 运行时Pid， Running、Stopping 时可用，启动时重置
	Pid int32 `json:"pid"`
	// 时间戳 秒， 启动、停止时设置
	Timestamp int64 `json:"timestamp"`
	// Exited 信息，启动时重置
	ExitMsg string `json:"exitMsg"`
	// 已经自动重启次数，启动时重置
	AutoStartTimes int `json:"autoStartTimes"`

	Cpu float64 `json:"cpu"`
	Mem float64 `json:"mem"`
}

type Process struct {
	ID       int              `json:"id"`
	Name     string           `json:"name"`
	Dir      string           `json:"dir"`
	Config   []*ProcessConfig `json:"config"`
	Command  string           `json:"command"`
	Labels   []string         `json:"labels"`
	Node     string           `json:"node"`
	User     string           `json:"user"`
	CreateAt int64            `json:"createAt"`

	State ProcessState `json:"state"`

	// 检测器
	MonitorState *MonitorState `json:"-"`

	// 子进程启动关闭优先级，优先级低的，最先启动，关闭的时候最后关闭
	// 默认值为999 。。非必须设置
	Priority int `json:"priority"`
	// 子进程启动多少秒之后，此时状态如果是running，则我们认为启动成功了
	// 默认值为2 。。非必须设置
	StartSecs int64 `json:"startSecs"`
	// 这个是当我们向子进程发送stop信号后，到系统返回信息所等待的最大时间。
	// 超过这个时间会向该子进程发送一个强制kill的信号。
	// 默认为10秒。。非必须设置
	StopWaitSecs int64 `json:"stopWaitSecs"`
	// 进程状态为 Exited时，自动重启
	// 默认为3次。。非必须设置
	AutoStartTimes int `json:"autoStartTimes"`
}

type ProcessMgr struct {
	GenID     int                 `json:"genId"`
	Process   map[int]*Process    `json:"process"`
	Monitor   *Monitor            `json:"monitor"`
	TagLabels map[string]struct{} `json:"-"`
	TagNodes  map[string]struct{} `json:"-"`
}

func (mgr *ProcessMgr) refreshLabels() {
	labels := map[string]struct{}{}
	nodes := map[string]struct{}{}
	for _, v := range mgr.Process {
		if _, ok := nodes[v.Node]; !ok {
			nodes[v.Node] = struct{}{}
		}

		for _, label := range v.Labels {
			if _, ok := labels[label]; !ok {
				labels[label] = struct{}{}
			}
		}
	}

	mgr.TagLabels = labels
	mgr.TagNodes = nodes

	if mgr.Monitor == nil {
		mgr.Monitor = &Monitor{
			Cpu:           1,
			Mem:           1,
			Disk:          0,
			Interval:      10,
			AlertInterval: 60,
		}
	}
}

func processTick() {
	rpcReq := map[string]*protocol.ProcessStateReq{}
	for _, p := range processMgr.Process {
		if !(p.State.Status == types.ProcStatusStopped ||
			p.State.Status == types.ProcStatusExited) {
			req, ok := rpcReq[p.Node]
			if !ok {
				req = &protocol.ProcessStateReq{
					Ids: make([]int32, 0, 4),
				}
				rpcReq[p.Node] = req
			}
			req.Ids = append(req.Ids, int32(p.ID))
		}
	}

	for n, req := range rpcReq {
		node, ok := nodeMgr.Nodes[n]
		if !ok || !node.Online() {
			// 节点不在线 设置状态为 unknown
			change := false
			for _, id := range req.GetIds() {
				p, ok := processMgr.Process[int(id)]
				if !ok {
					continue
				}
				if p.State.Status != types.ProcStatusUnknown {
					p.State.Status = types.ProcStatusUnknown
					change = true
				}
			}
			if change {
				saveStore(snProcessMgr)
			}
			continue
		}
		_ = center.Go(node, req, drpc.DefaultRPCTimeout, func(i interface{}, e error) {
			if e != nil {
				return
			}
			change := false
			rpcResp := i.(*protocol.ProcessStateResp)
			//log.Println(22, rpcResp)
			for id, state := range rpcResp.GetStates() {
				p, ok := processMgr.Process[int(id)]
				if !ok {
					continue
				}

				p.State.Pid = state.Pid
				p.State.Mem = state.GetMem()
				p.State.Cpu = state.GetCpu()
				status := types.ProcStatus(state.GetStatus())
				switch p.State.Status {
				case types.ProcStatusUnknown:
					p.State.Status = status
					switch status {
					case types.ProcStatusRunning:
					case types.ProcStatusStopped:
					case types.ProcStatusExited:
						p.State.AutoStartTimes = p.AutoStartTimes // 未知状态不重启
						p.State.ExitMsg = state.GetExitMsg()
					}
					change = true
				case types.ProcStatusStarting:
					switch status {
					case types.ProcStatusRunning:
						if NowUnix() >= p.State.Timestamp+p.StartSecs {
							// 启动时间
							p.State.Status = types.ProcStatusRunning
							change = true
						}
					case types.ProcStatusStopped:
						p.State.Status = types.ProcStatusStopped
						change = true
					case types.ProcStatusExited:
						p.State.Status = types.ProcStatusExited
						p.State.AutoStartTimes = p.AutoStartTimes // 启动阶段不重启
						p.State.ExitMsg = state.GetExitMsg()
						change = true
					}
				case types.ProcStatusRunning:
					switch status {
					case types.ProcStatusRunning:
						// 仅运行状态监控报警
						p.monitor(state.GetCpu(), state.GetMem())
					case types.ProcStatusStopped:
						p.State.Status = types.ProcStatusStopped
						change = true
					case types.ProcStatusExited:
						p.State.Status = types.ProcStatusExited
						p.State.ExitMsg = state.GetExitMsg()
						change = true
					}
				case types.ProcStatusStopping:
					switch status {
					case types.ProcStatusRunning:
						if NowUnix() >= p.State.Timestamp+p.StopWaitSecs {
							// 停止时间超时 ，强行停止
							_ = center.Go(node, &protocol.ProcessSignalReq{
								Pid:    p.State.Pid,
								Signal: int32(syscall.SIGKILL),
							}, drpc.DefaultRPCTimeout, func(i interface{}, e error) {})
						}
					case types.ProcStatusStopped:
						p.State.Status = types.ProcStatusStopped
						change = true
					case types.ProcStatusExited:
						p.State.Status = types.ProcStatusExited
						p.State.AutoStartTimes = p.AutoStartTimes // 停止阶段不重启
						p.State.ExitMsg = state.GetExitMsg()
						change = true
					}
				}
			}
			if change {
				saveStore(snProcessMgr)
			}
		})
	}
}

func processAutoStart() {
	for _, p := range processMgr.Process {
		if p.State.Status == types.ProcStatusExited &&
			p.State.AutoStartTimes < p.AutoStartTimes {

			node, ok := nodeMgr.Nodes[p.Node]
			if !ok || !node.Online() {
				continue
			}

			log.Printf("process %d auto start times %d\n", p.ID, p.State.AutoStartTimes)
			if err := p.start(node, func(code string, err error) {}); err == nil {
				p.State.AutoStartTimes += 1
				p.State.Status = types.ProcStatusStarting
				p.State.Timestamp = NowUnix()
				p.State.ExitMsg = ""
				saveStore(snProcessMgr)
			}
		}
	}
}

func (p *Process) start(node *Node, callback func(code string, err error)) error {
	configs := make(map[string]string, len(p.Config))
	for _, cfg := range p.Config {
		configs[cfg.Name] = cfg.Context
	}

	cmd := strings.ReplaceAll(p.Command, "{{path}}", fmt.Sprintf("%s/%s", common.Dir, p.Name))
	rpcReq := &protocol.ProcessExecReq{
		Id:     int32(p.ID),
		Dir:    p.Dir,
		Name:   p.Name,
		Args:   strings.Fields(cmd),
		Config: configs,
	}

	return center.Go(node, rpcReq, drpc.DefaultRPCTimeout, func(i interface{}, e error) {
		if e != nil {
			callback("", e)
			return
		}
		rpcResp := i.(*protocol.ProcessExecResp)
		callback(rpcResp.GetCode(), nil)
	})
}

func (p *Process) monitor(cpu, mem float64) {
	if processMgr.Monitor.Opened {
		if p.MonitorState == nil {
			p.MonitorState = new(MonitorState)
		}

		processMgr.Monitor.Alert(p.MonitorState, cpu, mem, 0, func() string {
			return fmt.Sprintf("应用名:%s|所属节点:%s", p.Name, p.Node)
		})

	} else {
		if p.MonitorState != nil {
			p.MonitorState = nil
		}
	}
}
