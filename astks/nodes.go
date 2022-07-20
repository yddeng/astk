package astks

import (
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/yddeng/astk/pkg/codec"
	"github.com/yddeng/astk/pkg/common"
	"github.com/yddeng/astk/pkg/protocol"
	"github.com/yddeng/dnet"
	"github.com/yddeng/dnet/drpc"
	"log"
	"net"
	"time"
)

type Center struct {
	token     string
	acceptor  dnet.Acceptor
	rpcServer *drpc.Server
	rpcClient *drpc.Client
}

func newCenter(address, token string) *Center {
	c := new(Center)
	c.token = token
	c.acceptor = dnet.NewTCPAcceptor(address)
	c.rpcClient = drpc.NewClient()
	c.rpcServer = drpc.NewServer()
	c.rpcServer.Register(proto.MessageName(&protocol.LoginReq{}), c.onLogin)
	log.Printf("tcp server run %s.\n", address)
	return c
}

func (c *Center) Go(n *node, data proto.Message, timeout time.Duration, callback func(interface{}, error)) error {
	return c.rpcClient.Go(n, proto.MessageName(data), data, timeout, callback)
}

func (c *Center) startListener() error {
	return c.acceptor.ServeFunc(func(conn net.Conn) {
		dnet.NewTCPSession(conn,
			dnet.WithCodec(new(codec.Codec)),
			dnet.WithTimeout(common.HeartbeatTimeout, 0),
			dnet.WithErrorCallback(func(session dnet.Session, err error) {
				log.Println(err)
				session.Close(err)
			}),
			dnet.WithMessageCallback(func(session dnet.Session, data interface{}) {
				taskQueue.Submit(func() {
					var err error
					switch data.(type) {
					case *drpc.Request:
						err = c.rpcServer.OnRPCRequest(&node{session: session}, data.(*drpc.Request))
					case *drpc.Response:
						err = c.rpcClient.OnRPCResponse(data.(*drpc.Response))
					case *codec.Message:
						c.dispatchMsg(session, data.(*codec.Message))
					}
					if err != nil {
						log.Println(err)
					}
				})
			}),
			dnet.WithCloseCallback(func(session dnet.Session, reason error) {
				taskQueue.Submit(func() {
					log.Printf("session closed, reason: %s\n", reason)
					ctx := session.Context()
					if ctx != nil {
						client := ctx.(*node)
						client.session = nil
						session.SetContext(nil)
					}
				})
			}))
	})

}

func (c *Center) dispatchMsg(session dnet.Session, msg *codec.Message) {
	cmd := msg.GetCmd()
	switch cmd {
	case codec.CmdHeartbeat:
		_ = session.Send(msg)
	case codec.CmdNodeState:
		ctx := session.Context()
		if ctx != nil {
			node := ctx.(*node)
			node.onNodeState(msg.GetData().(*protocol.NodeState))
		}
	default:

	}

}

type node struct {
	Name    string       `json:"name"`
	Inet    string       `json:"inet"`
	Net     string       `json:"net"`
	LoginAt int64        `json:"login_at"` // 登陆时间
	session dnet.Session `json:"_"`

	nodeState *protocol.NodeState `json:"_"`
}

func (n *node) Online() bool {
	return n.session != nil
}

func (n *node) SendRequest(req *drpc.Request) error {
	if n.session == nil {
		return errors.New("session is nil")
	}
	return n.session.Send(req)
}

func (n *node) SendResponse(resp *drpc.Response) error {
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
		channel.(*node).session.Close(errors.New("token failed. "))
		return
	}

	name := msg.GetName()
	client := nodes[name]
	if client == nil {
		client = &node{Name: name}
		nodes[name] = client
	}
	if client.session != nil {
		replier.Reply(&protocol.LoginResp{Code: "client already login. "}, nil)
		channel.(*node).session.Close(errors.New("client already login. "))
		return
	}

	client.Inet = msg.GetInet()
	client.Net = msg.GetNet()
	client.LoginAt = NowUnix()

	client.session = channel.(*node).session
	client.session.SetContext(client)
	log.Printf("onLogin %s", client.session.RemoteAddr().String())
	replier.Reply(&protocol.LoginResp{}, nil)
	saveStore(snNode)
}

func (n *node) onNodeState(msg *protocol.NodeState) {
	n.nodeState = msg
}
