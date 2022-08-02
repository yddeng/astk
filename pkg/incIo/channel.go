package incIo

import (
	"github.com/yddeng/dnet"
	"net"
)

type Channel struct {
	ChanID  int32
	Conn    net.Conn
	Session dnet.Session
}

func (this *Channel) HandRead(close func()) {

}

func (this *Channel) Write(close func()) {

}

func (this *Channel) Close() {

}
