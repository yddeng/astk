package astks

import (
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/yddeng/astk/pkg/util"
	"github.com/yddeng/dnet/dhttp"
	"testing"
	"time"
)

var (
	//address = "212.129.131.27:40156"
	address  = "10.128.2.123:40156"
	startWeb = false
)

func startWebListener(t *testing.T) {
	if !startWeb {
		return
	}

	var err error
	var cfg Config
	if err = util.DecodeJsonFromFile(&cfg, "../cmd/amps_config.json"); err != nil {
		t.Fatal(err)
	}

	if err = Start(cfg); err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Millisecond * 100)
}

func authLogin(t *testing.T, Username, Password string) string {
	req, _ := dhttp.PostJson(fmt.Sprintf("http://%s/auth/login", address), struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{Username: Username, Password: Password})

	ret, err := req.ToString()
	if err != nil {
		t.Fatal(err)
	}
	return ret
}

func TestAuthHandler_Login(t *testing.T) {
	startWebListener(t)

	ret := authLogin(t, "admin", "123456")
	t.Log(ret)
}

func TestAuthHandler_Logout(t *testing.T) {
	startWebListener(t)

	ret := authLogin(t, "admin", "123456")
	t.Log(ret, gjson.Get(ret, "data.token").String())

	req2, _ := dhttp.NewRequest(fmt.Sprintf("http://%s/auth/logout", address), "POST")
	req2.SetHeader("Access-Token", gjson.Get(ret, "data.token").String())

	ret, err := req2.ToString()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)

}

func TestRedirect(t *testing.T) {
	startWebListener(t)

	req, _ := dhttp.PostJson(fmt.Sprintf("http://%s/api/auth/login", address), struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{Username: "admin", Password: "123456"})

	ret, err := req.ToString()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}
