package astks

import (
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/yddeng/astk/pkg/codec"
	"github.com/yddeng/astk/pkg/common"
	"github.com/yddeng/astk/pkg/protocol"
	"github.com/yddeng/dnet"
	"github.com/yddeng/dnet/drpc"
	"github.com/yddeng/utils/task"
	"log"
	"net"
	"net/http"
	"os"
	"runtime"
	"time"
)

func NowUnix() int64 {
	return time.Now().Unix()
}

type Config struct {
	DataPath     string        `json:"data_path"`
	Ip           string        `json:"ip"`
	CenterConfig *CenterConfig `json:"center_config"`
	WebConfig    *WebConfig    `json:"web_config"`
}

type CenterConfig struct {
	Port  int    `json:"port"`
	Token string `json:"token"`
}

type WebConfig struct {
	Port     int    `json:"port"`
	Index    string `json:"index"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var (
	taskQueue = task.NewTaskPool(1, runtime.NumCPU()*1024)
	center    *Center
	app       *gin.Engine
	config    *Config
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
	return c
}

func (c *Center) Go(n *Node, data proto.Message, timeout time.Duration, callback func(interface{}, error)) error {
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
						err = c.rpcServer.OnRPCRequest(&Node{session: session}, data.(*drpc.Request))
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
						client := ctx.(*Node)
						client.session = nil
						session.SetContext(nil)
					}
				})
			}))
	})

}

func (c *Center) dispatchMsg(session dnet.Session, msg *codec.Message) {
	ctx := session.Context()
	if ctx == nil {
		return
	}
	node := ctx.(*Node)

	cmd := msg.GetCmd()
	switch cmd {
	case codec.CmdHeartbeat:
		_ = session.Send(msg)
	case codec.CmdNodeState:
		node.onNodeState(msg.GetData().(*protocol.NodeState))
	case codec.CmdChannelMessage:
		node.onChannelMessage(msg.GetData().(*protocol.ChannelMessage))
	default:

	}

}

func Start(cfg Config) (err error) {
	_ = os.MkdirAll(cfg.DataPath, os.ModePerm)
	if err = loadStore(cfg.DataPath); err != nil {
		return
	}
	config = &cfg

	{
		centerCfg := cfg.CenterConfig
		log.Printf("center server run %s:%d.\n", cfg.Ip, centerCfg.Port)
		center = newCenter(fmt.Sprintf("0.0.0.0:%d", centerCfg.Port), centerCfg.Token)
		go func() {
			if err := center.startListener(); err != nil {
				panic(err)
			}
		}()
	}

	{
		webCfg := cfg.WebConfig
		/*
		 所有的公共变量在队列中执行。
		 使用warp函数处理过的方法，已经是在队列中执行。
		*/

		app = gin.New()
		app.Use(gin.Logger(), gin.Recovery())

		// 前端
		if webCfg.Index != "" {
			app.Use(static.Serve("/", static.LocalFile(webCfg.Index, false)))
			app.NoRoute(func(ctx *gin.Context) {
				ctx.File(webCfg.Index + "/index.html")
			})
		}

		// 跨域
		app.Use(func(ctx *gin.Context) {
			ctx.Header("Access-Control-Allow-Origin", "*")
			ctx.Header("Access-Control-Allow-Headers", "*")
			ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH")
			ctx.Header("Access-Control-Allow-Credentials", "true")
			ctx.Header("Access-Control-Expose-Headers", "*")
			if ctx.Request.Method == "OPTIONS" {
				// 处理浏览器的options请求时，返回200状态即可
				ctx.JSON(http.StatusOK, "")
				ctx.Abort()
				return
			}

			ctx.Next()
		})

		initHandler(app)

		log.Printf("web server run %s:%d.\n", cfg.Ip, webCfg.Port)
		go func() {
			if err := app.Run(fmt.Sprintf("0.0.0.0:%d", webCfg.Port)); err != nil {
				panic(err)
			}
		}()

	}

	go func() {
		timer := time.NewTimer(time.Second)
		for {
			<-timer.C
			taskQueue.Submit(func() {
				processTick()
				processAutoStart()
				doSave(false)
				timer.Reset(time.Second)
			})
		}
	}()
	return nil
}

func Stop() {

}
