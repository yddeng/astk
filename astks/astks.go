package astks

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/yddeng/utils/task"
	"log"
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
	CenterConfig *CenterConfig `json:"center_config"`
	WebConfig    *WebConfig    `json:"web_config"`
}

type CenterConfig struct {
	Address string `json:"address"`
	Token   string `json:"token"`
}

type WebConfig struct {
	Address string `json:"address"`
	Index   string `json:"index"`
	Admin   struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"admin"`
}

var (
	taskQueue = task.NewTaskPool(1, runtime.NumCPU()*1024)
	center    *Center
	app       *gin.Engine
	webCfg    *WebConfig
)

func Start(cfg Config) (err error) {
	_ = os.MkdirAll(cfg.DataPath, os.ModePerm)
	if err = loadStore(cfg.DataPath); err != nil {
		return
	}

	centerCfg := cfg.CenterConfig
	log.Printf("center server run %s.\n", centerCfg.Address)
	center = newCenter(centerCfg.Address, centerCfg.Token)
	go func() {
		if err := center.startListener(); err != nil {
			panic(err)
		}
	}()

	webServiceRun(cfg.WebConfig)

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

func webServiceRun(cfg *WebConfig) {
	webCfg = cfg
	/*
	 所有的公共变量在队列中执行。
	 使用warp函数处理过的方法，已经是在队列中执行。
	*/

	app = gin.New()
	app.Use(gin.Logger(), gin.Recovery())

	// 前端
	if cfg.Index != "" {
		app.Use(static.Serve("/", static.LocalFile(cfg.Index, false)))
		app.NoRoute(func(ctx *gin.Context) {
			ctx.File(cfg.Index + "/index.html")
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

	log.Printf("web server run %s.\n", cfg.Address)
	go func() {
		if err := app.Run(cfg.Address); err != nil {
			panic(err)
		}
	}()
}
