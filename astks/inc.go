package astks

import (
	"fmt"
	"github.com/yddeng/astk/pkg/inc"
	"github.com/yddeng/astk/pkg/incIo"
	"github.com/yddeng/astk/pkg/protocol"
	"github.com/yddeng/dnet/drpc"
	"log"
	"net"
	"strconv"
	"time"
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
	RemotePort string   `json:"remotePort"` // 访问端口
	Node       string  `json:"node"`       // 被代理的节点
	LocalIP    string  `json:"localIp"`    // 节点被代理地址
	LocalPort  string   `json:"localPort"`  // 节点被代理端口
	Opened     bool    `json:"opened"`     // 是否开启

	listener net.Listener
	OpenID   int32                    `json:"-"`
	Channel  map[int32]*incIo.Channel `json:"-"`
}

type IncMgr struct {
	IncMap map[string]*Inc `json:"incMap"`
}

func (this *Inc) startListener() error {
	switch this.Type {
	case IncTypeHttp, IncTypeHttps, IncTypeTCP:
		return this.startTcpListener()
	default:
		return nil
	}
}

func (this *Inc) startTcpListener() (err error) {
	this.listener, err = net.Listen("tcp", net.JoinHostPort("0.0.0.0", strconv.Itoa(int(this.RemotePort))))
	if err != nil {
		return
	}

	go func() {
		for {
			c, err := this.listener.Accept()
			if err != nil {
				if ne, ok := err.(net.Error); ok && ne.Temporary() {
					time.Sleep(time.Millisecond * 5)
					continue
				}
			}

			go this.openChannel(c)
		}
	}()
	return
}

func (this *Inc) openChannel(conn net.Conn) {
	srcIp,srcPort,err := net.SplitHostPort(conn.RemoteAddr().String())
	if err != nil{
		log.Println(err)
	}
	req := &protocol.OpenConnectionReq{
		Id:                   this.ID,
		OpenID:               this.OpenID,
		SrcIp:                ,
		SrcPort:              0,
	}
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
