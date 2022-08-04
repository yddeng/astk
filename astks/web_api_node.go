package astks

import (
	"github.com/yddeng/astk/pkg/protocol"
	"log"
	"sort"
)

type nodeInfo struct {
	Name    string              `json:"name"`
	Inet    string              `json:"inet"`
	Net     string              `json:"net"`
	LoginAt int64               `json:"loginAt"` // 登陆时间
	Online  bool                `json:"online"`
	State   *protocol.NodeState `json:"state"`
}

type nodeHandler struct{}

func (*nodeHandler) List(wait *WaitConn, user string, req struct {
	PageNo   int `json:"pageNo"`
	PageSize int `json:"pageSize"`
}) {
	//log.Printf("%s by(%s) %v\n", done.route, user, req)

	s := make([]*nodeInfo, 0, len(nodeMgr.Nodes))
	for _, n := range nodeMgr.Nodes {
		s = append(s, &nodeInfo{
			Name:    n.Name,
			Inet:    n.Inet,
			Net:     n.Net,
			LoginAt: n.LoginAt,
			Online:  n.Online(),
			State:   n.nodeState,
		})
	}
	sort.Slice(s, func(i, j int) bool {
		if s[i].Online == s[j].Online {
			return s[i].LoginAt > s[j].LoginAt
		} else {
			return s[i].Online
		}
	})

	start, end := listRange(req.PageNo, req.PageSize, len(nodeMgr.Nodes))
	wait.SetResult("", pageData{
		PageNo:     req.PageNo,
		PageSize:   req.PageSize,
		TotalCount: len(s),
		Data:       s[start:end],
	})
	wait.Done()
}

func (*nodeHandler) Status(wait *WaitConn, user string) {
	log.Printf("%s by(%s)\n", wait.route, user)
	type result struct {
		All    []string `json:"all"`
		Online []string `json:"online"`
	}

	ret := &result{
		All:    make([]string, 0, len(nodeMgr.Nodes)),
		Online: make([]string, 0, len(nodeMgr.Nodes)),
	}
	for _, n := range nodeMgr.Nodes {
		ret.All = append(ret.All, n.Name)
		if n.Online() {
			ret.Online = append(ret.Online, n.Name)
		}
	}
	wait.SetResult("", ret)
	wait.Done()
}

func (*nodeHandler) Remove(wait *WaitConn, user string, req struct {
	Name string `json:"name"`
}) {
	log.Printf("%s by(%s) %v\n", wait.route, user, req)
	defer func() { wait.Done() }()
	n, ok := nodeMgr.Nodes[req.Name]
	if !ok || n.Online() {
		wait.SetResult("当前状态不允许移除", nil)
		return
	}

	delete(nodeMgr.Nodes, req.Name)
	saveStore(snNodeMgr)
}

func (*nodeHandler) Bell(wait *WaitConn, user string, req struct {
	Name string `json:"name"`
	Bell bool   `json:"bell"`
}) {
	// log.Printf("%s by(%s) %v\n", wait.route, user, req)
	defer func() { wait.Done() }()

	n, ok := nodeMgr.Nodes[req.Name]
	if !ok {
		wait.SetResult("不存在的节点", nil)
		return
	}

	if n.Bell != req.Bell {
		n.Bell = req.Bell
		saveStore(snNodeMgr)
	}
}
