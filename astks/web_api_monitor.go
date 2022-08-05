package astks

import (
	"log"
)

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

func (*monitorHandler) Update(wait *WaitConn, user string, req struct {
	Type    string   `json:"type"`
	Monitor *Monitor `json:"monitor"`
}) {
	log.Printf("%s by(%s) %v\n", wait.route, user, req)

	if req.Type == "node" {
		nodeMgr.Monitor = req.Monitor
		saveStore(snNodeMgr)
	} else if req.Type == "process" {
		processMgr.Monitor = req.Monitor
		saveStore(snProcessMgr)
	} else {
		wait.SetResult("类型错误", nil)
	}

	wait.Done()
}
