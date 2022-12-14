package astks

import (
	"github.com/yddeng/astk/pkg/protocol"
	"github.com/yddeng/astk/pkg/types"
	"github.com/yddeng/dnet/drpc"
	"log"
	"sort"
	"syscall"
)

type processHandler struct{}

func (*processHandler) Tags(wait *WaitConn, user string) {
	//log.Printf("%s by(%s) \n", wait .route, user)
	defer func() { wait.Done() }()

	wait.SetResult("", struct {
		Nodes  map[string]struct{} `json:"nodes"`
		Labels map[string]struct{} `json:"labels"`
	}{Nodes: processMgr.TagNodes, Labels: processMgr.TagLabels})
}

func (*processHandler) findProcess(Nodes, Labels map[string]struct{}, Status map[types.ProcStatus]struct{}) []*Process {
	s := make([]*Process, 0, len(processMgr.Process))
	for _, v := range processMgr.Process {
		if len(Nodes) > 0 {
			if _, ok := Nodes[v.Node]; !ok {
				continue
			}
		}
		if len(Labels) > 0 {
			hasLabel := false
			for _, label := range v.Labels {
				if _, ok := Labels[label]; ok {
					hasLabel = true
					break
				}
			}
			if !hasLabel {
				continue
			}
		}

		if len(Status) > 0 {
			if _, ok := Status[v.State.Status]; !ok {
				continue
			}
		}

		s = append(s, v)
	}
	return s
}

func (this *processHandler) List(wait *WaitConn, user string, req struct {
	Nodes    map[string]struct{}           `json:"nodes"`
	Labels   map[string]struct{}           `json:"labels"`
	Status   map[types.ProcStatus]struct{} `json:"status"`
	PageNo   int                           `json:"pageNo"`
	PageSize int                           `json:"pageSize"`
}) {
	//log.Printf("%s by(%s) %v\n", wait.route, user, req)
	defer func() { wait.Done() }()

	s := this.findProcess(req.Nodes, req.Labels, req.Status)
	sort.Slice(s, func(i, j int) bool {
		return s[i].ID < s[j].ID
	})
	start, end := listRange(req.PageNo, req.PageSize, len(s))
	wait.SetResult("", pageData{
		PageNo:     req.PageNo,
		PageSize:   req.PageSize,
		TotalCount: len(s),
		Data:       s[start:end],
	})
}

func (this *processHandler) Create(wait *WaitConn, user string, req struct {
	Name           string           `json:"name"`
	Dir            string           `json:"dir"`
	Config         []*ProcessConfig `json:"config"`
	Command        string           `json:"command"`
	Labels         []string         `json:"labels"`
	Node           string           `json:"node"`
	Priority       int              `json:"priority"`
	StartSecs      int64            `json:"startSecs"`
	StopWaitSecs   int64            `json:"stopWaitSecs"`
	AutoStartTimes int              `json:"autoStartTimes"`
}) {
	log.Printf("%s by(%s) %v\n", wait.route, user, req)
	defer func() { wait.Done() }()

	for _, p := range processMgr.Process {
		if p.Name == req.Name {
			wait.SetResult("???????????????", nil)
			return
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
		Status: types.ProcStatusStopped,
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
	processMgr.refreshLabels()
	saveStore(snProcessMgr)
}

func (this *processHandler) Update(wait *WaitConn, user string, req struct {
	ID             int              `json:"id"`
	Name           string           `json:"name"`
	Dir            string           `json:"dir"`
	Config         []*ProcessConfig `json:"config"`
	Command        string           `json:"command"`
	Labels         []string         `json:"labels"`
	Node           string           `json:"node"`
	Priority       int              `json:"priority"`
	StartSecs      int64            `json:"startSecs"`
	StopWaitSecs   int64            `json:"stopWaitSecs"`
	AutoStartTimes int              `json:"autoStartTimes"`
}) {
	log.Printf("%s by(%s) %v\n", wait.route, user, req)
	defer func() { wait.Done() }()

	for _, p := range processMgr.Process {
		if p.Name == req.Name && p.ID != req.ID {
			wait.SetResult("???????????????", nil)
			return
		}
	}

	p, ok := processMgr.Process[req.ID]
	if !ok || !(p.State.Status == types.ProcStatusStopped ||
		p.State.Status == types.ProcStatusExited) {
		wait.SetResult("???????????????????????????", nil)
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

	processMgr.refreshLabels()
	saveStore(snProcessMgr)
}

func (*processHandler) Delete(wait *WaitConn, user string, req struct {
	ID int `json:"id"`
}) {
	log.Printf("%s by(%s) %v\n", wait.route, user, req)
	defer func() { wait.Done() }()

	p, ok := processMgr.Process[req.ID]
	if !ok || !(p.State.Status == types.ProcStatusStopped ||
		p.State.Status == types.ProcStatusExited) {
		wait.SetResult("???????????????????????????", nil)
		return
	}

	delete(processMgr.Process, req.ID)
	processMgr.refreshLabels()
	saveStore(snProcessMgr)
}

func (*processHandler) Start(wait *WaitConn, user string, req struct {
	ID int `json:"id"`
}) {
	log.Printf("%s by(%s) %v\n", wait.route, user, req)
	p, ok := processMgr.Process[req.ID]
	if !ok ||
		!(p.State.Status == types.ProcStatusStopped ||
			p.State.Status == types.ProcStatusExited) {
		wait.SetResult("???????????????????????????", nil)
		wait.Done()
		return
	}

	node, ok := nodeMgr.Nodes[p.Node]
	if !ok || !node.Online() {
		wait.SetResult("???????????????", nil)
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
			Status:    types.ProcStatusStarting,
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
	if !ok || p.State.Status != types.ProcStatusRunning {
		wait.SetResult("???????????????????????????", nil)
		wait.Done()
		return
	}

	node, ok := nodeMgr.Nodes[p.Node]
	if !ok || !node.Online() {
		wait.SetResult("???????????????", nil)
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
		p.State.Status = types.ProcStatusStopping
		p.State.Timestamp = NowUnix()
		saveStore(snProcessMgr)
	}
}

func (this *processHandler) BatchStart(wait *WaitConn, user string, req struct {
	Nodes  map[string]struct{}           `json:"nodes"`
	Labels map[string]struct{}           `json:"labels"`
	Status map[types.ProcStatus]struct{} `json:"status"`
}) {
	log.Printf("%s by(%s) %v\n", wait.route, user, req)

	req.Status = map[types.ProcStatus]struct{}{
		types.ProcStatusExited:  {},
		types.ProcStatusStopped: {},
	}
	s := this.findProcess(req.Nodes, req.Labels, req.Status)
	// ???????????????????????????
	sort.Slice(s, func(i, j int) bool {
		return s[i].Priority < s[j].Priority
	})
	for _, p := range s {
		node, ok := nodeMgr.Nodes[p.Node]
		if !ok || !node.Online() {
			continue
		}

		if err := p.start(node, func(code string, err error) {}); err == nil {
			p.State = ProcessState{
				Status:    types.ProcStatusStarting,
				Timestamp: NowUnix(),
			}
			saveStore(snProcessMgr)
		}
	}

	wait.Done()
}

func (this *processHandler) BatchStop(wait *WaitConn, user string, req struct {
	Nodes  map[string]struct{}           `json:"nodes"`
	Labels map[string]struct{}           `json:"labels"`
	Status map[types.ProcStatus]struct{} `json:"status"`
}) {
	log.Printf("%s by(%s) %v\n", wait.route, user, req)

	req.Status = map[types.ProcStatus]struct{}{
		types.ProcStatusRunning: {},
	}
	s := this.findProcess(req.Nodes, req.Labels, req.Status)
	// ???????????????????????????
	sort.Slice(s, func(i, j int) bool {
		return s[i].Priority > s[j].Priority
	})
	for _, p := range s {
		node, ok := nodeMgr.Nodes[p.Node]
		if !ok || !node.Online() {
			continue
		}

		rpcReq := &protocol.ProcessSignalReq{
			Pid:    p.State.Pid,
			Signal: int32(syscall.SIGTERM),
		}

		if err := center.Go(node, rpcReq, drpc.DefaultRPCTimeout, func(i interface{}, e error) {}); err == nil {
			p.State.Status = types.ProcStatusStopping
			p.State.Timestamp = NowUnix()
			saveStore(snProcessMgr)
		}
	}

	wait.Done()
}

func (*processHandler) TailLog(wait *WaitConn, user string, req struct {
	ID    int `json:"id"`
	Start int `json:"start"`
}) {
	// log.Printf("%s by(%s) %v\n", wait.route, user, req)

	p, ok := processMgr.Process[req.ID]
	if !ok || !(p.State.Status == types.ProcStatusRunning || p.State.Status == types.ProcStatusStarting || p.State.Status == types.ProcStatusStopping) {
		wait.SetResult("?????????????????????", nil)
		wait.Done()
		return
	}

	node, ok := nodeMgr.Nodes[p.Node]
	if !ok || !node.Online() {
		wait.SetResult("???????????????", nil)
		wait.Done()
		return
	}

	rpcReq := &protocol.TailLogReq{
		Id:    int32(req.ID),
		Start: int32(req.Start),
	}

	if err := center.Go(node, rpcReq, drpc.DefaultRPCTimeout, func(i interface{}, e error) {
		if e != nil {
			wait.Done()
			return
		}
		rpcResp := i.(*protocol.TailLogResp)
		wait.SetResult("", struct {
			Context string `json:"context"`
			End     int32  `json:"end"`
		}{Context: string(rpcResp.GetContext()), End: rpcResp.GetEnd()})

		wait.Done()
	}); err != nil {
		wait.SetResult(err.Error(), nil)
		wait.Done()
	}
}
