package astks

import (
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/yddeng/astk/pkg/types"
	"github.com/yddeng/dnet/dhttp"
	"testing"
)

func TestIncHandler_Create(t *testing.T) {
	startWebListener(t)

	ret := authLogin(t, "admin", "123456")
	t.Log(ret, gjson.Get(ret, "data.token").String())

	req2, _ := dhttp.NewRequest(fmt.Sprintf("http://%s/inc/create", address), "POST")
	req2.SetHeader("Access-Token", gjson.Get(ret, "data.token").String())
	req2, _ = req2.WriteJSON(struct {
		Type       types.IncType `json:"type"`
		RemotePort string        `json:"remotePort"` // 访问端口
		Node       string        `json:"node"`       // 被代理的节点
		LocalIP    string        `json:"localIp"`    // 节点被代理地址
		LocalPort  string        `json:"localPort"`  // 节点被代理端口
	}{
		Type:       types.IncTypeTCP,
		RemotePort: "22335",
		Node:       "astke",
		LocalIP:    "10.128.2.123",
		LocalPort:  "22",
	})

	ret, err := req2.ToString()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)

}

func TestIncHandler_List(t *testing.T) {
	startWebListener(t)

	ret := authLogin(t, "admin", "123456")
	t.Log(ret, gjson.Get(ret, "data.token").String())

	req2, _ := dhttp.NewRequest(fmt.Sprintf("http://%s/inc/list", address), "POST")
	req2.SetHeader("Access-Token", gjson.Get(ret, "data.token").String())
	req2, _ = req2.WriteJSON(struct {
		PageNo   int `json:"pageNo"`
		PageSize int `json:"pageSize"`
	}{PageNo: 1, PageSize: 10})

	ret, err := req2.ToString()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}

func TestIncHandler_Opened(t *testing.T) {
	startWebListener(t)

	ret := authLogin(t, "admin", "123456")
	t.Log(ret, gjson.Get(ret, "data.token").String())

	req2, _ := dhttp.NewRequest(fmt.Sprintf("http://%s/inc/opened", address), "POST")
	req2.SetHeader("Access-Token", gjson.Get(ret, "data.token").String())
	req2, _ = req2.WriteJSON(struct {
		ID     string `json:"id"`
		Opened bool   `json:"opened"`
	}{ID: "NUfIQDMavsxBRoHkmE3Q2O", Opened: true})

	ret, err := req2.ToString()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)
}
