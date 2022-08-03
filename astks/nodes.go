package astks

import (
	"errors"
	"fmt"
	"github.com/yddeng/astk/pkg/protocol"
	"github.com/yddeng/dnet"
	"github.com/yddeng/dnet/drpc"
	"log"
	"strconv"
)

type NodeMgr struct {
	Nodes   map[string]*Node `json:"nodes"`
	Monitor *Monitor         `json:"monitor"`
}

type Node struct {
	Name    string `json:"name"`
	Inet    string `json:"inet"`
	Net     string `json:"net"`
	LoginAt int64  `json:"loginAt"` // 登陆时间

	session dnet.Session

	nodeState *protocol.NodeState

	// 检测器
	Bell         bool          `json:"bell"`
	MonitorState *MonitorState `json:"-"`
}

func (n *Node) Online() bool {
	return n.session != nil
}

func (n *Node) SendRequest(req *drpc.Request) error {
	if n.session == nil {
		return errors.New("session is nil")
	}
	return n.session.Send(req)
}

func (n *Node) SendResponse(resp *drpc.Response) error {
	if n.session == nil {
		return errors.New("session is nil")
	}
	return n.session.Send(resp)
}

func (c *Center) onLogin(replier *drpc.Replier, req interface{}) {
	channel := replier.Channel
	msg := req.(*protocol.LoginReq)
	log.Printf("onLogin %v", msg)

	if c.token != "" && msg.GetToken() != c.token {
		replier.Reply(&protocol.LoginResp{Code: "token failed"}, nil)
		channel.(*Node).session.Close(errors.New("token failed. "))
		return
	}

	name := msg.GetName()
	client := nodeMgr.Nodes[name]
	if client == nil {
		client = &Node{Name: name, Bell: true}
		nodeMgr.Nodes[name] = client
	}
	if client.session != nil {
		replier.Reply(&protocol.LoginResp{Code: "client already login. "}, nil)
		channel.(*Node).session.Close(errors.New("client already login. "))
		return
	}

	client.Inet = msg.GetInet()
	client.Net = msg.GetNet()
	client.LoginAt = NowUnix()

	client.session = channel.(*Node).session
	client.session.SetContext(client)
	log.Printf("onLogin %s", client.session.RemoteAddr().String())
	replier.Reply(&protocol.LoginResp{}, nil)
	saveStore(snNodeMgr)
}

func (n *Node) onNodeState(msg *protocol.NodeState) {
	n.nodeState = msg

	cpuUsed, _ := strconv.ParseFloat(msg.GetCpu()["usedPercent"], 64)
	memUsed, _ := strconv.ParseFloat(msg.GetMem()["virtualUsedPercent"], 64)
	diskUsed, _ := strconv.ParseFloat(msg.GetDisk()["usedPercent"], 64)
	n.monitor(cpuUsed, memUsed, diskUsed)
}

func (n *Node) monitor(cpu, mem, disk float64) {
	if n.Bell {
		if n.MonitorState == nil {
			n.MonitorState = new(MonitorState)
		}

		nodeMgr.Monitor.Alert(n.MonitorState, cpu, mem, disk, func() string {
			return fmt.Sprintf("节点名:%s", n.Name)
		})

	} else {
		if n.MonitorState != nil {
			n.MonitorState = nil
		}
	}
}
