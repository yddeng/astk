package astke

import (
	"fmt"
	"github.com/yddeng/astk/pkg/incIo"
	"github.com/yddeng/astk/pkg/protocol"
	"github.com/yddeng/astk/pkg/types"
	"github.com/yddeng/dnet/drpc"
	"log"
	"net"
)

type Channel struct {
	ID             string
	ChanID         int32
	Type           types.IncType
	IP, Port       string
	SrcIP, SrcPort string

	ioc *incIo.Channel
}

var channels = map[int32]*Channel{}

func (this *Channel) dial() (net.Conn, error) {
	switch this.Type {
	case types.IncTypeHttp, types.IncTypeHttps, types.IncTypeTCP:
		return net.Dial("tcp", net.JoinHostPort(this.IP, this.Port))
	default:
		return nil, fmt.Errorf("invaild type %s", this.Type)
	}
}

func (this *Channel) close(chanID int32) {
	er.Submit(func() {
		//log.Printf("channel %d closed", this.ChanID)
		delete(channels, this.ChanID)
	})
}

func (er *Executor) onOpenChannel(replier *drpc.Replier, req interface{}) {
	msg := req.(*protocol.OpenChannelReq)
	log.Printf("onOpenChannel %v", msg)

	chanID := msg.GetChanID()
	if channel, ok := channels[chanID]; ok {
		channel.ioc.Close()
	}

	channel := &Channel{
		ID:      msg.GetId(),
		ChanID:  chanID,
		Type:    types.IncType(msg.GetType()),
		IP:      msg.GetIp(),
		Port:    msg.GetPort(),
		SrcIP:   msg.GetSrcIp(),
		SrcPort: msg.GetSrcPort(),
	}

	conn, err := channel.dial()
	if err != nil {
		_ = replier.Reply(&protocol.OpenChannelResp{Code: err.Error()}, nil)
		return
	}

	// test
	if chanID == 0 {
		_ = conn.Close()
	} else {
		channel.ioc = incIo.New(conn, er.session, channel.close)
		channel.ioc.ChanID = channel.ChanID
		channel.ioc.ID = channel.ID
		go channel.ioc.ChanReader()
		go channel.ioc.ChanWriter()
		channels[chanID] = channel
	}

	_ = replier.Reply(&protocol.OpenChannelResp{}, nil)
}

func (er *Executor) onChannelMessage(msg *protocol.ChannelMessage) {
	chanID := msg.GetChanID()
	channel, ok := channels[chanID]
	if !ok {
		return
	}

	if msg.GetEof() {
		channel.ioc.Close()
	} else {
		channel.ioc.Write(msg)
	}

}
