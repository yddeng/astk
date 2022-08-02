package astke

import (
	"fmt"
	"github.com/yddeng/astk/pkg/incIo"
	"github.com/yddeng/astk/pkg/protocol"
	"github.com/yddeng/dnet/drpc"
	"net"
)

type Inc struct {
	ID   string
	Type string
	IP   string
	Port int32

	Channel map[int32]*incIo.Channel
}

func (this *Inc) dial() (net.Conn, error) {
	return net.Dial("tcp", net.JoinHostPort(this.IP, fmt.Sprintf("%d", this.Port)))
}

var incMap = map[string]*Inc{}

func (er *Executor) onCreateDialer(replier *drpc.Replier, req interface{}) {
	msg := req.(*protocol.CreateDialerReq)
	// log.Printf("onTailLog %v", msg)

	if inc, ok := incMap[msg.GetId()]; ok {
		for _, channel := range inc.Channel {
			channel.Close()
		}
		delete(incMap, msg.GetId())
	}

	inc := &Inc{
		ID:      msg.GetId(),
		Type:    msg.GetType(),
		IP:      msg.GetIp(),
		Port:    msg.GetPort(),
		Channel: map[int32]*incIo.Channel{},
	}

	// test
	conn, err := inc.dial()
	if err != nil {
		_ = replier.Reply(&protocol.CreateDialerResp{Code: err.Error()}, nil)
		return
	}
	_ = conn.Close()

	incMap[msg.GetId()] = inc

	_ = replier.Reply(&protocol.CreateDialerResp{}, nil)
}

func (er *Executor) onOpenChannel(replier *drpc.Replier, req interface{}) {
	msg := req.(*protocol.OpenConnectionReq)
	// log.Printf("onTailLog %v", msg)

	inc, ok := incMap[msg.GetId()]
	if !ok {
		_ = replier.Reply(&protocol.OpenConnectionResp{Code: "not exist"}, nil)
		return
	}

	if channel, ok := inc.Channel[msg.GetOpenID()]; ok {
		channel.Close()
		delete(inc.Channel, msg.GetOpenID())
	}

	conn, err := inc.dial()
	if err != nil {
		_ = replier.Reply(&protocol.OpenConnectionResp{Code: err.Error()}, nil)
		return
	}

	channel := &incIo.Channel{
		ChanID:  msg.GetOpenID(),
		Conn:    conn,
		Session: er.session,
	}

	_ = replier.Reply(&protocol.CreateDialerResp{}, nil)
}
