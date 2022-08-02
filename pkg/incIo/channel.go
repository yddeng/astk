package incIo

import "net"

type Channel struct {
	ChanID int32
	Conn   net.Conn
}

func (this *Channel) HandRead(close func()) {

}

func (this *Channel) Write(close func()) {

}

func (this *Channel) Close() {

}
