package astks

import (
	"amp/back-go/protocol"
	"fmt"
	"github.com/yddeng/dnet/drpc"
	"log"
	"sort"
	"strings"
	"time"
)

type cmdHandler struct {
}

func (*cmdHandler) List(wait *WaitConn, user string, req struct {
	PageNo   int `json:"pageNo"`
	PageSize int `json:"pageSize"`
}) {
	//log.Printf("%s by(%s) %v\n", wait.route, user, req)
	defer func() { wait.Done() }()

	s := make([]*Cmd, 0, len(cmdMgr.CmdMap))
	for _, v := range cmdMgr.CmdMap {
		s = append(s, v)
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i].CallNo > s[j].CallNo
	})

	start, end := listRange(req.PageNo, req.PageSize, len(s))
	wait.SetResult("", struct {
		PageNo     int    `json:"pageNo"`
		PageSize   int    `json:"pageSize"`
		TotalCount int    `json:"totalCount"`
		Success    int    `json:"success"`
		Failed     int    `json:"failed"`
		CmdList    []*Cmd `json:"cmdList"`
	}{PageNo: req.PageNo,
		PageSize:   req.PageSize,
		TotalCount: len(s),
		Success:    cmdMgr.Success,
		Failed:     cmdMgr.Failed,
		CmdList:    s[start:end],
	})
}

func (*cmdHandler) Create(wait *WaitConn, user string, req struct {
	Name    string            `json:"name"`
	Dir     string            `json:"dir"`
	Context string            `json:"context"`
	Args    map[string]string `json:"args"`
}) {
	log.Printf("%s by(%s) %v\n", wait.route, user, req)
	defer func() { wait.Done() }()

	for _, cmd := range cmdMgr.CmdMap {
		if cmd.Name == req.Name {
			wait.SetResult("名字重复", nil)
			return
		}
	}

	if len(cmdContextReg(req.Context)) != len(req.Args) {
		wait.SetResult("变量与默认值数量不一致", nil)
		return
	}

	nowUnix := NowUnix()
	cmdMgr.GenID++
	id := cmdMgr.GenID
	cmd := &Cmd{
		ID:       id,
		Name:     req.Name,
		Dir:      req.Dir,
		Context:  req.Context,
		Args:     req.Args,
		User:     user,
		UpdateAt: nowUnix,
		CreateAt: nowUnix,
	}

	cmdMgr.CmdMap[id] = cmd
	saveStore(snCmdMgr)
}

func (*cmdHandler) Delete(wait *WaitConn, user string, req struct {
	ID int `json:"id"`
}) {
	log.Printf("%s by(%s) %v\n", wait.route, user, req)
	defer func() { wait.Done() }()

	if _, ok := cmdMgr.CmdMap[req.ID]; !ok {
		wait.SetResult("不存在的命令", nil)
		return
	}
	delete(cmdMgr.CmdMap, req.ID)
	delete(cmdMgr.CmdLogs, req.ID)
	saveStore(snCmdMgr)
}

func (*cmdHandler) Update(wait *WaitConn, user string, req struct {
	ID      int               `json:"id"`
	Name    string            `json:"name"`
	Dir     string            `json:"dir"`
	Context string            `json:"context"`
	Args    map[string]string `json:"args"`
}) {
	log.Printf("%s by(%s) %v\n", wait.route, user, req)
	defer func() { wait.Done() }()

	for _, cmd := range cmdMgr.CmdMap {
		if cmd.Name == req.Name && cmd.ID != req.ID {
			wait.SetResult("命令名重复", nil)
			return
		}
	}

	if cmd, ok := cmdMgr.CmdMap[req.ID]; !ok {
		wait.SetResult("不存在的命令", nil)
	} else {
		if len(cmdContextReg(req.Context)) != len(req.Args) {
			wait.SetResult("变量与默认值数量不一致", nil)
			return
		}

		cmd.Name = req.Name
		cmd.Dir = req.Dir
		cmd.Context = req.Context
		cmd.Args = req.Args
		cmd.User = user
		cmd.UpdateAt = NowUnix()
		saveStore(snCmdMgr)
	}
}

const (
	cmdDefaultTimeout = 60
	cmdMinTimeout     = 10
	cmdMaxTimeout     = 86400
)

func (*cmdHandler) Exec(wait *WaitConn, user string, req struct {
	ID      int               `json:"id"`
	Dir     string            `json:"dir"`
	Args    map[string]string `json:"args"`
	Node    string            `json:"node"`
	Timeout int               `json:"timeout"`
}) {
	log.Printf("%s by(%s) %v\n", wait.route, user, req)

	cmd, ok := cmdMgr.CmdMap[req.ID]
	if !ok {
		wait.SetResult("不存在的命令", nil)
		wait.Done()
		return
	}

	node, ok := nodes[req.Node]
	if !ok || !node.Online() {
		wait.SetResult("执行客户端不存在或不在线", nil)
		wait.Done()
		return
	}

	if cmd.doing == nil {
		cmd.doing = map[string]struct{}{}
	}
	if _, ok := cmd.doing[req.Node]; ok {
		wait.SetResult("当前命令正在该节点上执行", nil)
		wait.Done()
		return
	}

	context := cmd.Context
	for k, v := range req.Args {
		context = strings.ReplaceAll(context, fmt.Sprintf("{{%s}}", k), v)
	}

	if len(cmdContextReg(context)) > 0 {
		wait.SetResult("命令中存在未赋值变量", nil)
		wait.Done()
		return
	}

	// 超时时间
	if req.Timeout <= 0 {
		req.Timeout = cmdDefaultTimeout
	} else if req.Timeout < cmdMinTimeout {
		req.Timeout = cmdMinTimeout
	} else if req.Timeout > cmdMaxTimeout {
		req.Timeout = cmdMaxTimeout
	}

	// 执行日志

	cmdLog := &CmdLog{
		CreateAt: NowUnix(),
		Timeout:  req.Timeout,
		User:     user,
		Node:     req.Node,
		Dir:      req.Dir,
		Context:  context,
	}

	cmdResult := func(cmdLog *CmdLog, ok bool, ret string) {
		if ok {
			cmdMgr.Success += 1
		} else {
			cmdMgr.Failed += 1
		}
		cmdLog.ResultAt = NowUnix()
		cmdLog.Result = ret
		saveStore(snCmdMgr)
	}

	rpcReq := &protocol.CmdExecReq{
		Dir:     req.Dir,
		Name:    "/bin/sh",
		Args:    []string{"-c", context},
		Timeout: int32(req.Timeout),
	}
	timeout := time.Second*time.Duration(req.Timeout) + drpc.DefaultRPCTimeout
	if err := center.Go(node, rpcReq, timeout, func(i interface{}, e error) {
		if e != nil {
			wait.SetResult(e.Error(), nil)
			cmdResult(cmdLog, false, e.Error())
			wait.Done()
			return
		}
		rpcResp := i.(*protocol.CmdExecResp)
		if rpcResp.GetCode() != "" {
			wait.SetResult(rpcResp.GetCode(), nil)
			cmdResult(cmdLog, false, rpcResp.GetCode())
		} else {
			cmdResult(cmdLog, true, rpcResp.GetOutStr())
			wait.SetResult("", cmdLog)
		}
		delete(cmd.doing, req.Node)
		wait.Done()
	}); err != nil {
		log.Println(err)
		wait.SetResult(err.Error(), nil)
		wait.Done()
	} else {
		cmd.doing[req.Node] = struct{}{}
		cmd.CallNo++
		cmdLog.ID = cmd.CallNo
		cmdMgr.CmdLogs[req.ID] = append([]*CmdLog{cmdLog}, cmdMgr.CmdLogs[req.ID]...)
		if len(cmdMgr.CmdLogs[req.ID]) > cmdLogCapacity {
			cmdMgr.CmdLogs[req.ID] = cmdMgr.CmdLogs[req.ID][:cmdLogCapacity]
		}
		saveStore(snCmdMgr)
	}

}

func (*cmdHandler) Log(wait *WaitConn, user string, req struct {
	ID       int `json:"id"`
	PageNo   int `json:"pageNo"`
	PageSize int `json:"pageSize"`
}) {
	//log.Printf("%s by(%s) %v\n", wait.route, user, req)
	defer func() { wait.Done() }()

	logs, ok := cmdMgr.CmdLogs[req.ID]
	if !ok {
		logs = []*CmdLog{}
	}

	start, end := listRange(req.PageNo, req.PageSize, len(logs))
	wait.SetResult("", struct {
		PageNo     int       `json:"pageNo"`
		PageSize   int       `json:"pageSize"`
		TotalCount int       `json:"totalCount"`
		LogList    []*CmdLog `json:"logList"`
	}{PageNo: req.PageNo,
		PageSize:   req.PageSize,
		TotalCount: len(logs),
		LogList:    logs[start:end],
	})

}
