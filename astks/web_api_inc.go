package astks

import (
	token2 "github.com/yddeng/astk/pkg/token"
	"github.com/yddeng/astk/pkg/types"
	"log"
	"sort"
)

type incHandler struct {
}

func (*incHandler) List(wait *WaitConn, user string, req struct {
	PageNo   int `json:"pageNo"`
	PageSize int `json:"pageSize"`
}) {
	//log.Printf("%s by(%s) %v\n", done.route, user, req)

	type Item struct {
		*Inc
		Opened  bool `json:"opened"`
		Channel int  `json:"channel"`
	}

	s := make([]*Item, 0, len(incMgr.IncMap))
	for _, n := range incMgr.IncMap {
		s = append(s, &Item{
			Inc:     n,
			Opened:  n.opened(),
			Channel: len(n.channels),
		})
	}
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
	wait.Done()
}

func (*incHandler) Create(wait *WaitConn, user string, req struct {
	Type       types.IncType `json:"type"`
	RemotePort string        `json:"remotePort"` // 访问端口
	Node       string        `json:"node"`       // 被代理的节点
	LocalIP    string        `json:"localIp"`    // 节点被代理地址
	LocalPort  string        `json:"localPort"`  // 节点被代理端口
	Desc       string        `json:"desc"`
}) {
	log.Printf("%s by(%s) %v\n", wait.route, user, req)
	defer func() { wait.Done() }()

	switch req.Type {
	case types.IncTypeTCP, types.IncTypeHttp, types.IncTypeHttps:
	default:
		wait.SetResult("类型错误", nil)
		return
	}

	inc := &Inc{
		Type:       req.Type,
		RemotePort: req.RemotePort,
		Node:       req.Node,
		LocalIP:    req.LocalIP,
		LocalPort:  req.LocalPort,
		Desc:       req.Desc,
	}

	id := token2.GenToken(22)
	for {
		if _, ok := incMgr.IncMap[id]; !ok {
			break
		} else {
			id = token2.GenToken(22)
		}
	}

	inc.ID = id
	incMgr.IncMap[id] = inc
	saveStore(snIncMgr)
}

func (*incHandler) Delete(wait *WaitConn, user string, req struct {
	ID string `json:"id"`
}) {
	log.Printf("%s by(%s) %v\n", wait.route, user, req)
	defer func() { wait.Done() }()

	inc, ok := incMgr.IncMap[req.ID]
	if !ok {
		wait.SetResult("操作对像不存在", nil)
		return
	}

	inc.stop()
	delete(incMgr.IncMap, req.ID)
	saveStore(snIncMgr)

}

func (*incHandler) Opened(wait *WaitConn, user string, req struct {
	ID     string `json:"id"`
	Opened bool   `json:"opened"`
}) {
	log.Printf("%s by(%s) %v\n", wait.route, user, req)

	inc, ok := incMgr.IncMap[req.ID]
	if !ok {
		wait.SetResult("操作对像不存在", nil)
		wait.Done()
		return
	}

	if req.Opened == inc.opened() {
		wait.Done()
		return
	}

	if req.Opened {
		inc.start(func(err string) {
			if err != "" {
				wait.SetResult(err, nil)
			}
			wait.Done()
		})
	} else {
		inc.stop()
		wait.Done()
	}

}
