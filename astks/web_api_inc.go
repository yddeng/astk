package astks

import (
	inc2 "github.com/yddeng/astk/pkg/inc"
	token2 "github.com/yddeng/astk/pkg/token"
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

	s := make([]*Inc, 0, len(incMgr.IncMap))
	for _, n := range incMgr.IncMap {
		s = append(s, n)
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
	Type       IncType `json:"type"`
	RemotePort int32   `json:"remotePort"` // 访问端口
	Node       string  `json:"node"`       // 被代理的节点
	LocalIP    string  `json:"localIp"`    // 节点被代理地址
	LocalPort  int32   `json:"localPort"`  // 节点被代理端口
}) {
	log.Printf("%s by(%s) %v\n", wait.route, user, req)

	inc := &Inc{
		Type:       req.Type,
		RemotePort: req.RemotePort,
		Node:       req.Node,
		LocalIP:    req.LocalIP,
		LocalPort:  req.LocalPort,
		Opened:     false,
	}

	id := token2.GenToken(22)
	for {
		if _, ok := incMgr.IncMap[id]; !ok {
			break
		} else {
			id = token2.GenToken(22)
		}
	}

	incMgr.IncMap[id] = inc
	saveStore(snIncMgr)
	wait.Done()
}

func (*incHandler) Delete(wait *WaitConn, user string, req struct {
	ID string `json:"id"`
}) {
	log.Printf("%s by(%s) %v\n", wait.route, user, req)

	inc, ok := incMgr.IncMap[req.ID]
	if !ok {
		wait.SetResult("操作对像不存在", nil)
		wait.Done()
		return
	}

	inc.closeDialer()
	inc.Opened = false
	inc.Channel = map[int32]*inc2.Channel{}

	saveStore(snIncMgr)
	wait.Done()

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

	if req.Opened == inc.Opened {
		wait.Done()
		return
	}

	if req.Opened {
		if err := inc.createDialer(func(err string) {
			if err != "" {
				wait.SetResult(err, nil)
			} else {
				inc.Opened = true
				saveStore(snIncMgr)
			}
			wait.Done()
		}); err != nil {
			wait.SetResult(err.Error(), nil)
			wait.Done()
			return
		}
	} else {
		inc.closeDialer()
		inc.Opened = false
		inc.Channel = map[int32]*inc2.Channel{}
		saveStore(snIncMgr)
		wait.Done()
	}

}
