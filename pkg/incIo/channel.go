package incIo

import (
	"github.com/yddeng/astk/pkg/codec"
	"github.com/yddeng/astk/pkg/protocol"
	"github.com/yddeng/dnet"
	"net"
	"sync"
)

type Channel struct {
	ID        string
	ChanID    int32
	conn      net.Conn
	session   dnet.Session
	sendCh    chan *protocol.ChannelMessage
	closeOnce sync.Once
	closeCh   chan bool
	closeFn   func(chanID int32)
}

func New(conn net.Conn, session dnet.Session, closeFn func(chanID int32)) *Channel {
	return &Channel{
		conn:    conn,
		session: session,
		sendCh:  make(chan *protocol.ChannelMessage, 1024),
		closeCh: make(chan bool),
		closeFn: closeFn,
	}
}

func (this *Channel) Close() {
	this.closeOnce.Do(func() {
		if this.conn != nil {
			_ = this.conn.Close()
		}

		close(this.closeCh)

		_ = this.session.Send(codec.NewMessage(&protocol.ChannelMessage{
			Id:     this.ID,
			ChanID: this.ChanID,
			Eof:    true,
		}))

		this.closeFn(this.ChanID)
	})
}

func (this *Channel) Write(msg *protocol.ChannelMessage) {
	this.sendCh <- msg
}

func (this *Channel) ChanWriter() {
	defer func() {
		//log.Printf("channel %d write closed", this.ChanID)
		this.Close()
	}()

	for {
		select {
		case msg := <-this.sendCh:
			if _, err := this.conn.Write(msg.GetData()); err != nil {
				//log.Printf("channel %d conn write error %s", this.ChanID, err)
				return
			}
		case <-this.closeCh:
			return
		}

	}
}

func (this *Channel) ChanReader() {
	defer func() {
		//log.Printf("channel %d read closed", this.ChanID)
		this.Close()
	}()

	for {
		buf := make([]byte, 2*1024)
		n, err := this.conn.Read(buf)
		if err != nil {
			//log.Printf("channel %d conn read error %s", this.ChanID, err)
			return
		}

		msg := &protocol.ChannelMessage{
			Id:     this.ID,
			ChanID: this.ChanID,
			Data:   buf[:n],
		}
		if err = this.session.Send(codec.NewMessage(msg)); err != nil {
			//log.Printf("channel %d session send error %s", this.ChanID, err)
			return
		}
	}
}
