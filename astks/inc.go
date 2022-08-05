package astks

import (
	"fmt"
	"github.com/yddeng/astk/pkg/incIo"
	"github.com/yddeng/astk/pkg/protocol"
	"github.com/yddeng/astk/pkg/types"
	"github.com/yddeng/dnet/drpc"
	"log"
	"net"
	"time"
)

type Inc struct {
	ID         string        `json:"id"`
	Type       types.IncType `json:"type"`
	RemotePort string        `json:"remotePort"` // 访问端口
	Node       string        `json:"node"`       // 被代理的节点
	LocalIP    string        `json:"localIp"`    // 节点被代理地址
	LocalPort  string        `json:"localPort"`  // 节点被代理端口
	Desc       string        `json:"desc"`

	listener net.Listener
	channels map[int32]*incIo.Channel
}

func (this *Inc) opened() bool {
	return this.listener != nil
}

func (this *Inc) close(chanID int32) {
	taskQueue.Submit(func() {
		//log.Printf("channel %d closed", chanID)
		delete(this.channels, chanID)
	})
}

type IncMgr struct {
	IncMap map[string]*Inc `json:"incMap"`
	genID  int32
}

func (this *Inc) startListener() error {
	switch this.Type {
	case types.IncTypeHttp, types.IncTypeHttps, types.IncTypeTCP:
		return this.startTcpListener()
	default:
		return fmt.Errorf("类型未实现")
	}
}

func (this *Inc) startTcpListener() (err error) {
	this.listener, err = net.Listen("tcp", net.JoinHostPort("0.0.0.0", this.RemotePort))
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
				taskQueue.Submit(func() { this.stop() })
				return
			}

			taskQueue.Submit(func() {
				node, ok := nodeMgr.Nodes[this.Node]
				if !ok || !node.Online() {
					log.Printf("节点%s不存在或不在线", this.Node)
					c.Close()
					this.stop()
					return
				}

				incMgr.genID++
				this.openChannel(node, c, incMgr.genID)
			})

		}
	}()
	return
}

func (this *Inc) start(callback func(err string)) {
	if this.listener == nil {
		node, ok := nodeMgr.Nodes[this.Node]
		if !ok || !node.Online() {
			callback("节点离线或不存在")
			return
		}

		req := &protocol.OpenChannelReq{
			Id:   this.ID,
			Type: string(this.Type),
			Ip:   this.LocalIP,
			Port: this.LocalPort,
		}
		center.Go(node, req, drpc.DefaultRPCTimeout, func(i interface{}, e error) {
			if e != nil {
				callback(e.Error())
				return
			}
			rpcResp := i.(*protocol.OpenChannelResp)
			if rpcResp.GetCode() != "" {
				callback(rpcResp.GetCode())
				return
			}

			if err := this.startListener(); err != nil {
				callback(err.Error())
				return
			}

			this.channels = map[int32]*incIo.Channel{}

			callback("")
		})
	} else {
		callback("")
	}
}

func (this *Inc) stop() {
	if this.listener != nil {
		_ = this.listener.Close()
		this.listener = nil
	}

	for _, channel := range this.channels {
		channel.Close()
	}
}

func (this *Inc) openChannel(node *Node, conn net.Conn, chanID int32) {
	srcIp, srcPort, err := net.SplitHostPort(conn.RemoteAddr().String())
	if err != nil {
		log.Println(err)
	}

	req := &protocol.OpenChannelReq{
		Id:      this.ID,
		ChanID:  chanID,
		Type:    string(this.Type),
		Ip:      this.LocalIP,
		Port:    this.LocalPort,
		SrcIp:   srcIp,
		SrcPort: srcPort,
	}
	center.Go(node, req, drpc.DefaultRPCTimeout, func(i interface{}, e error) {
		if e != nil {
			log.Println(err)
			conn.Close()
			return
		}
		rpcResp := i.(*protocol.OpenChannelResp)
		if rpcResp.GetCode() != "" {
			log.Println(err)
			conn.Close()
			return
		}

		ioc := incIo.New(conn, node.session, this.close)
		ioc.ChanID = chanID
		ioc.ID = this.ID
		go ioc.ChanReader()
		go ioc.ChanWriter()
		this.channels[chanID] = ioc
	})
}

func (n *Node) onChannelMessage(msg *protocol.ChannelMessage) {
	id := msg.GetId()
	chanID := msg.GetChanID()

	if inc, ok := incMgr.IncMap[id]; ok {
		if channel, ok := inc.channels[chanID]; ok {
			if msg.GetEof() {
				channel.Close()
			} else {
				channel.Write(msg)
			}
		}
	}

}
