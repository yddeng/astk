package astks

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"sync"
)

// 应答结构
type Result struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type WaitConn struct {
	code     int
	ctx      *gin.Context
	route    string
	result   Result
	done     chan struct{}
	doneOnce sync.Once
}

func newWaitConn(ctx *gin.Context, route string) *WaitConn {
	return &WaitConn{
		ctx:   ctx,
		code:  http.StatusOK,
		route: route,
		done:  make(chan struct{}),
	}
}

func (this *WaitConn) Done(code ...int) {
	this.doneOnce.Do(func() {
		if this.result.Message == "" {
			this.result.Success = true
		}
		if len(code) > 0 {
			this.code = code[0]
		}
		close(this.done)
	})
}

func (this *WaitConn) Context() *gin.Context {
	return this.ctx
}

func (this *WaitConn) SetResult(message string, data interface{}) {
	this.result.Message = message
	this.result.Data = data
}

func (this *WaitConn) Wait() {
	<-this.done
}

type webTask func()

func (t webTask) Do() {
	t()
}

func transBegin(ctx *gin.Context, fn interface{}, args ...reflect.Value) {
	val := reflect.ValueOf(fn)
	if val.Kind() != reflect.Func {
		panic("value not func")
	}
	typ := val.Type()
	if typ.NumIn() != len(args)+2 {
		panic("func argument error")
	}

	route := getCurrentRoute(ctx)
	wait := newWaitConn(ctx, route)
	if err := taskQueue.SubmitTask(webTask(func() {
		user, ok := checkToken(ctx, route)
		if !ok {
			wait.SetResult("Token验证失败", nil)
			wait.Done(401)
			return
		}

		ok = checkPermission(ctx, route, user)
		if !ok {
			wait.SetResult("无操作权限", nil)
			wait.Done(403)
			return
		}
		val.Call(append([]reflect.Value{reflect.ValueOf(wait), reflect.ValueOf(user)}, args...))
	}), true); err != nil {
		wait.SetResult("访问人数过多", nil)
		wait.Done()
	}
	wait.Wait()

	ctx.JSON(wait.code, wait.result)
}

func getCurrentRoute(ctx *gin.Context) string {
	return ctx.FullPath()
}

func getJsonBody(ctx *gin.Context, inType reflect.Type) (inValue reflect.Value, err error) {
	if inType.Kind() == reflect.Ptr {
		inValue = reflect.New(inType.Elem())
	} else {
		inValue = reflect.New(inType)
	}
	if err = ctx.ShouldBindJSON(inValue.Interface()); err != nil {
		return
	}
	if inType.Kind() != reflect.Ptr {
		inValue = inValue.Elem()
	}
	return
}

func warpHandle(fn interface{}) gin.HandlerFunc {
	val := reflect.ValueOf(fn)
	if val.Kind() != reflect.Func {
		panic("value not func")
	}
	typ := val.Type()
	switch typ.NumIn() {
	case 2: // func(wait *WaitConn, username string)
		return func(ctx *gin.Context) {
			transBegin(ctx, fn)
		}
	case 3: // func(wait *WaitConn, username string,req struct)
		return func(ctx *gin.Context) {
			inValue, err := getJsonBody(ctx, typ.In(2))
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": "Json unmarshal failed!",
					"error":   err.Error(),
				})
				return
			}

			transBegin(ctx, fn, inValue)
		}
	default:
		panic("func symbol error")
	}
}

var (
	// 允许无token的路由
	allowTokenRoute = map[string]struct{}{
		"/auth/login":  {},
		"/auth/logout": {},
	}
	// 允许无权限的路由
	allowPermissionRoute = map[string]struct{}{
		"/auth/login":  {},
		"/auth/logout": {},
	}
)

func checkToken(ctx *gin.Context, route string) (user string, ok bool) {
	if _, ok = allowTokenRoute[route]; ok {
		return
	}
	tkn := ctx.GetHeader("Access-Token")
	if tkn == "" {
		ok = false
		return
	}
	if user, ok = getTknUser(tkn); !ok {
		ok = false
		return
	}
	ok = true
	return
}

func checkPermission(ctx *gin.Context, route, user string) (ok bool) {
	if _, ok = allowPermissionRoute[route]; ok {
		return
	}
	ok = true
	return
}

type pageData struct {
	PageNo     int         `json:"pageNo"`
	PageSize   int         `json:"pageSize"`
	TotalCount int         `json:"totalCount"`
	Data       interface{} `json:"dataList"`
}

func listRange(pageNo, pageSize, length int) (start int, end int) {
	start = (pageNo - 1) * pageSize
	if start < 0 {
		start = 0
	}
	if start > length {
		start = length
	}
	end = start + pageSize
	if end > length {
		end = length
	}
	return
}

func initHandler(app *gin.Engine) {
	authHandle := new(authHandler)
	authGroup := app.Group("/auth")
	authGroup.POST("/login", warpHandle(authHandle.Login))
	authGroup.POST("/logout", warpHandle(authHandle.Logout))

	nodeHandle := new(nodeHandler)
	nodeGroup := app.Group("/node")
	nodeGroup.POST("/list", warpHandle(nodeHandle.List))
	nodeGroup.POST("/names", warpHandle(nodeHandle.Names))
	nodeGroup.POST("/remove", warpHandle(nodeHandle.Remove))
	nodeGroup.POST("/bell", warpHandle(nodeHandle.Bell))

	cmdHandle := new(cmdHandler)
	cmdGroup := app.Group("/cmd")
	cmdGroup.POST("/list", warpHandle(cmdHandle.List))
	cmdGroup.POST("/create", warpHandle(cmdHandle.Create))
	cmdGroup.POST("/delete", warpHandle(cmdHandle.Delete))
	cmdGroup.POST("/update", warpHandle(cmdHandle.Update))
	cmdGroup.POST("/exec", warpHandle(cmdHandle.Exec))
	cmdGroup.POST("/log", warpHandle(cmdHandle.Log))

	processHandle := new(processHandler)
	processGroup := app.Group("/process")
	processGroup.POST("/tags", warpHandle(processHandle.Tags))
	processGroup.POST("/list", warpHandle(processHandle.List))
	processGroup.POST("/create", warpHandle(processHandle.Create))
	processGroup.POST("/delete", warpHandle(processHandle.Delete))
	processGroup.POST("/update", warpHandle(processHandle.Update))
	processGroup.POST("/start", warpHandle(processHandle.Start))
	processGroup.POST("/stop", warpHandle(processHandle.Stop))
	processGroup.POST("/batch/start", warpHandle(processHandle.BatchStart))
	processGroup.POST("/batch/stop", warpHandle(processHandle.BatchStop))
	processGroup.POST("/tail", warpHandle(processHandle.TailLog))
	processGroup.POST("/bell", warpHandle(processHandle.Bell))

	monitorHandle := new(monitorHandler)
	monitorGroup := app.Group("/monitor")
	monitorGroup.POST("/info", warpHandle(monitorHandle.Info))
	monitorGroup.POST("/rule", warpHandle(monitorHandle.SetRule))
	monitorGroup.POST("/notify", warpHandle(monitorHandle.SetNotify))
}
