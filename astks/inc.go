package astks

import (
	"fmt"
	"github.com/yddeng/astk/pkg/inc"
	"github.com/yddeng/astk/pkg/incIo"
	"github.com/yddeng/astk/pkg/protocol"
	"github.com/yddeng/dnet/drpc"
	"log"
)

type IncType string

const (
	IncTypeHttp  IncType = "http"
	IncTypeHttps IncType = "https"
	IncTypeTCP   IncType = "tcp"
)

type Inc struct {
	ID         string  `json:"id"`
	Type       IncType `json:"type"`
	RemotePort int32   `json:"remotePort"` // 访问端口
	Node       string  `json:"node"`       // 被代理的节点
	LocalIP    string  `json:"localIp"`    // 节点被代理地址
	LocalPort  int32   `json:"localPort"`  // 节点被代理端口
	Opened     bool    `json:"opened"`     // 是否开启

	OpenID  int32                    `json:"-"`
	Channel map[int32]*incIo.Channel `json:"-"`
}

type IncMgr struct {
	IncMap map[string]*Inc `json:"incMap"`
}

func (this *Inc) createDialer(callback func(err string)) error {
	node, ok := nodeMgr.Nodes[this.Node]
	if !ok || !node.Online() {
		callback(fmt.Sprintf("节点%s不存在或不在线", this.Node))
		return nil
	}

	req := &protocol.CreateDialerReq{
		Type: string(this.Type),
		Ip:   this.LocalIP,
		Port: this.LocalPort,
	}

	return center.Go(node, req, drpc.DefaultRPCTimeout, func(i interface{}, e error) {
		if e != nil {
			callback(e.Error())
			return
		}
		rpcResp := i.(*protocol.CreateDialerResp)
		callback(rpcResp.GetCode())
	})
}

func (this *Inc) closeDialer() {
	node, ok := nodeMgr.Nodes[this.Node]
	if !ok || !node.Online() {
		return
	}

	if len(this.Channel) > 0 {
		for id := range this.Channel {
			req := &protocol.CloseConnectionReq{
				OpenID: id,
			}
			center.Go(node, req, drpc.DefaultRPCTimeout, func(i interface{}, e error) {
				if e != nil {
					log.Println(e)
				}
			})
		}
	}
}
