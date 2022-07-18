package astks

import (
	"log"
)

var (
	webCfg *WebConfig
)

type authHandler struct {
}

func (*authHandler) Login(wait *WaitConn, user string, req struct {
	Username string `json:"username"`
	Password string `json:"password"`
}) {

	log.Printf("%s %v\n", wait.route, req)
	defer func() { wait.Done() }()

	if webCfg.Admin.Username != req.Username || webCfg.Admin.Password != req.Password {
		wait.SetResult("用户或密码错误", nil)
		return
	}

	token := addToken(req.Username)
	wait.SetResult("", struct {
		Token string `json:"token"`
	}{Token: token})
	return
}

func (*authHandler) Logout(wait *WaitConn, user string) {
	log.Printf("%s by(%s) \n", wait.route, user)
	defer func() { wait.Done() }()
	if user == "" {
		return
	}
	rmUserTkn(user)
}
