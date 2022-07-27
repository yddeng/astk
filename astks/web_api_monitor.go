package astks

import "log"

type monitorHandler struct{}

func (*monitorHandler) Info(wait *WaitConn, user string, req struct {
	Type string `json:"type"`
}) {

	if req.Type == "node" {
		wait.SetResult("", nodeMgr.Monitor)
	} else if req.Type == "process" {
		wait.SetResult("", processMgr.Monitor)
	} else {
		wait.SetResult("类型错误", nil)
	}

	wait.Done()
}

func (*monitorHandler) SetRule(wait *WaitConn, user string, req struct {
	Type          string `json:"type"`
	Cpu           int    `json:"cpu"`
	Mem           int    `json:"mem"`
	Disk          int    `json:"disk"`
	Interval      int64  `json:"interval"`
	AlertInterval int64  `json:"continuityInterval"`
}) {
	log.Printf("%s by(%s) %v\n", wait.route, user, req)

	if req.Type == "node" {
		nodeMgr.Monitor.Cpu = req.Cpu
		nodeMgr.Monitor.Mem = req.Mem
		nodeMgr.Monitor.Disk = req.Disk
		nodeMgr.Monitor.Interval = req.Interval
		nodeMgr.Monitor.AlertInterval = req.AlertInterval
		saveStore(snNodeMgr)
	} else if req.Type == "process" {
		processMgr.Monitor.Cpu = req.Cpu
		processMgr.Monitor.Mem = req.Mem
		processMgr.Monitor.Disk = req.Disk
		processMgr.Monitor.Interval = req.Interval
		processMgr.Monitor.AlertInterval = req.AlertInterval
		saveStore(snProcessMgr)
	} else {
		wait.SetResult("类型错误", nil)
	}

	wait.Done()
}

func (*monitorHandler) SetNotify(wait *WaitConn, user string, req struct {
	Type         string        `json:"type"`
	NotifyType   MsgNotifyType `json:"notifyType"`
	NotifyServer string        `json:"notifyServer"`
}) {
	log.Printf("%s by(%s) %v\n", wait.route, user, req)

	if req.Type == "node" {
		nodeMgr.Monitor.Notify = &Notify{
			NotifyType:   req.NotifyType,
			NotifyServer: req.NotifyServer,
		}
		saveStore(snNodeMgr)
	} else if req.Type == "process" {
		processMgr.Monitor.Notify = &Notify{
			NotifyType:   req.NotifyType,
			NotifyServer: req.NotifyServer,
		}
		saveStore(snProcessMgr)
	} else {
		wait.SetResult("类型错误", nil)
	}

	wait.Done()
}
