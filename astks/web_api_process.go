package astks

import (
	"amp/back-go/common"
	"amp/back-go/protocol"
	"github.com/yddeng/dnet/drpc"
	"log"
	"syscall"
)

type processHandler struct{}

func (*processHandler) Labels(wait *WaitConn, user string) {
	//log.Printf("%s by(%s) \n", wait .route, user)
	defer func() { wait.Done() }()
	wait.SetResult("", processMgr.Labels)
}

func (this *processHandler) LabelAdd(wait *WaitConn, user string, req struct {
	Label string `json:"label"`
}) {
	log.Printf("%s by(%s) %v\n", wait.route, user, req)
	defer func() { wait.Done() }()

	if _, ok := processMgr.Labels[req.Label]; !ok {
		processMgr.Labels[req.Label] = struct{}{}
		saveStore(snProcessMgr)
	}
}

func (*processHandler) LabelRemove(wait *WaitConn, user string, req struct {
	Label string `json:"label"`
}) {
	log.Printf("%s by(%s) %v\n", wait.route, user, req)
	defer func() { wait.Done() }()

	if _, ok := processMgr.Labels[req.Label]; !ok {
		wait.SetResult("不存在的分组", nil)
		return
	}

	for _, v := range processMgr.Process {
		if _, ok := v.Labels[req.Label]; ok {
			wait.SetResult("当前分组还存在进程，不允许删除", nil)
			return
		}
	}

	delete(processMgr.Labels, req.Label)
	saveStore(snProcessMgr)
}

func (*processHandler) List(wait *WaitConn, user string, req struct {
	Labels map[string]struct{} `json:"labels"`
}) {
	//log.Printf("%s by(%s) %v\n", wait .route, user, req)
	defer func() { wait.Done() }()

	if len(req.Labels) == 0 {
		wait.SetResult("", processMgr.Process)
	} else {
		s := make(map[int]*Process, len(processMgr.Process))
		for _, v := range processMgr.Process {
			for label := range v.Labels {
				if _, ok := req.Labels[label]; ok {
					s[v.ID] = v
					break
				}
			}
		}
		wait.SetResult("", s)
	}
}

func (this *processHandler) Create(wait *WaitConn, user string, req struct {
	Name           string              `json:"name"`
	Dir            string              `json:"dir"`
	Config         []*ProcessConfig    `json:"config"`
	Command        string              `json:"command"`
	Labels         map[string]struct{} `json:"labels"`
	Node           string              `json:"node"`
	Priority       int                 `json:"priority"`
	StartSecs      int64               `json:"startSecs"`
	StopWaitSecs   int64               `json:"stopWaitSecs"`
	AutoStartTimes int                 `json:"autoStartTimes"`
}) {
	log.Printf("%s by(%s) %v\n", wait.route, user, req)
	defer func() { wait.Done() }()

	for _, p := range processMgr.Process {
		if p.Name == req.Name {
			wait.SetResult("程序名重复", nil)
			return
		}
	}

	for labels := range req.Labels {
		if _, ok := processMgr.Labels[labels]; !ok {
			processMgr.Labels[labels] = struct{}{}
			saveStore(snProcessMgr)
		}
	}

	processMgr.GenID++
	id := processMgr.GenID
	p := new(Process)
	p.ID = id
	p.Name = req.Name
	p.Dir = req.Dir
	p.Config = req.Config
	p.Command = req.Command
	p.State = ProcessState{
		Status: common.StateStopped,
	}
	p.Labels = req.Labels
	p.Node = req.Node
	p.User = user
	p.CreateAt = NowUnix()
	p.Priority = req.Priority
	p.StartSecs = req.StartSecs
	p.StopWaitSecs = req.StopWaitSecs
	p.AutoStartTimes = req.AutoStartTimes

	processMgr.Process[id] = p
	saveStore(snProcessMgr)
}

func (this *processHandler) Update(wait *WaitConn, user string, req struct {
	ID             int                 `json:"id"`
	Name           string              `json:"name"`
	Dir            string              `json:"dir"`
	Config         []*ProcessConfig    `json:"config"`
	Command        string              `json:"command"`
	Labels         map[string]struct{} `json:"labels"`
	Node           string              `json:"node"`
	Priority       int                 `json:"priority"`
	StartSecs      int64               `json:"startSecs"`
	StopWaitSecs   int64               `json:"stopWaitSecs"`
	AutoStartTimes int                 `json:"autoStartTimes"`
}) {
	log.Printf("%s by(%s) %v\n", wait.route, user, req)
	defer func() { wait.Done() }()

	for _, p := range processMgr.Process {
		if p.Name == req.Name && p.ID != req.ID {
			wait.SetResult("程序名重复", nil)
			return
		}
	}

	for labels := range req.Labels {
		if _, ok := processMgr.Labels[labels]; !ok {
			processMgr.Labels[labels] = struct{}{}
			saveStore(snProcessMgr)
		}
	}

	p, ok := processMgr.Process[req.ID]
	if !ok || !(p.State.Status == common.StateStopped ||
		p.State.Status == common.StateExited) {
		wait.SetResult("当前状态不允许修改", nil)
		return
	}

	p.Name = req.Name
	p.Dir = req.Dir
	p.Config = req.Config
	p.Command = req.Command
	p.Labels = req.Labels
	p.Node = req.Node
	p.Priority = req.Priority
	p.StartSecs = req.StartSecs
	p.StopWaitSecs = req.StopWaitSecs
	p.AutoStartTimes = req.AutoStartTimes

	saveStore(snProcessMgr)
}

func (*processHandler) Delete(wait *WaitConn, user string, req struct {
	ID int `json:"id"`
}) {
	log.Printf("%s by(%s) %v\n", wait.route, user, req)
	defer func() { wait.Done() }()

	p, ok := processMgr.Process[req.ID]
	if !ok || !(p.State.Status == common.StateStopped ||
		p.State.Status == common.StateExited) {
		wait.SetResult("当前状态不允许删除", nil)
		return
	}

	delete(processMgr.Process, req.ID)
	saveStore(snProcessMgr)
}

func (*processHandler) Start(wait *WaitConn, user string, req struct {
	ID int `json:"id"`
}) {
	log.Printf("%s by(%s) %v\n", wait.route, user, req)
	p, ok := processMgr.Process[req.ID]
	if !ok ||
		!(p.State.Status == common.StateStopped ||
			p.State.Status == common.StateExited) {
		wait.SetResult("当前状态不允许启动", nil)
		wait.Done()
		return
	}

	node, ok := nodes[p.Node]
	if !ok || !node.Online() {
		wait.SetResult("节点无服务", nil)
		wait.Done()
		return
	}

	if err := p.start(node, func(code string, err error) {
		if err == nil {
			wait.SetResult(code, nil)
		}
		wait.Done()
	}); err != nil {
		wait.SetResult(err.Error(), nil)
		wait.Done()
	} else {
		p.State = ProcessState{
			Status:    common.StateStarting,
			Timestamp: NowUnix(),
		}
		saveStore(snProcessMgr)
	}
}

func (*processHandler) Stop(wait *WaitConn, user string, req struct {
	ID int `json:"id"`
}) {
	log.Printf("%s by(%s) %v\n", wait.route, user, req)

	p, ok := processMgr.Process[req.ID]
	if !ok || p.State.Status != common.StateRunning {
		wait.SetResult("当前状态不允许停止", nil)
		wait.Done()
		return
	}

	node, ok := nodes[p.Node]
	if !ok || !node.Online() {
		wait.SetResult("节点无服务", nil)
		wait.Done()
		return
	}

	rpcReq := &protocol.ProcessSignalReq{
		Pid:    p.State.Pid,
		Signal: int32(syscall.SIGTERM),
	}

	if err := center.Go(node, rpcReq, drpc.DefaultRPCTimeout, func(i interface{}, e error) {
		if e != nil {
			wait.Done()
			return
		}
		rpcResp := i.(*protocol.ProcessSignalResp)
		if rpcResp.GetCode() != "" {
			wait.SetResult(rpcResp.GetCode(), nil)
		}
		wait.Done()
	}); err != nil {
		wait.SetResult(err.Error(), nil)
		wait.Done()
	} else {
		p.State.Status = common.StateStopping
		p.State.Timestamp = NowUnix()
		saveStore(snProcessMgr)
	}
}
