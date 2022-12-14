package astks

import (
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/yddeng/dnet/dhttp"
	"regexp"
	"strings"
	"testing"
)

func TestCmdHandler_List(t *testing.T) {
	startWebListener(t)

	ret := authLogin(t, "admin", "123456")
	t.Log(ret, gjson.Get(ret, "data.token").String())

	token := gjson.Get(ret, "data.token").String()
	{
		//create
		shStr := `
set -euv
# 注释
sleep 11s
mkdir {{name}}
echo ok`

		req, _ := dhttp.NewRequest(fmt.Sprintf("http://%s/cmd/create", address), "POST")
		req.SetHeader("Access-Token", token)
		req.WriteJSON(struct {
			Name    string            `json:"name"`
			Dir     string            `json:"dir"`
			Context string            `json:"context"`
			Args    map[string]string `json:"args"`
		}{Name: "test", Dir: "", Context: shStr, Args: map[string]string{"name": "tttt"}})

		ret, err := req.ToString()
		t.Log(err, ret)
	}

	{
		//list
		req, _ := dhttp.NewRequest(fmt.Sprintf("http://%s/cmd/list", address), "POST")
		req.SetHeader("Access-Token", token)
		req.WriteJSON(struct {
			PageNo   int `json:"pageNo"`
			PageSize int `json:"pageSize"`
		}{PageNo: 1, PageSize: 10})

		ret, err := req.ToString()
		t.Log(err, ret)
	}

	{
		// exec
		req, _ := dhttp.NewRequest(fmt.Sprintf("http://%s/cmd/exec", address), "POST")
		req.SetHeader("Access-Token", token)
		req.WriteJSON(struct {
			Name    string            `json:"name"`
			Dir     string            `json:"dir"`
			Args    map[string]string `json:"args"`
			Node    string            `json:"node"`
			Timeout int               `json:"timeout"`
		}{Name: "test", Dir: "", Args: map[string]string{"name": "tttt"}, Node: "executor", Timeout: 12})

		ret, err := req.ToString()
		t.Log(err, ret)
		out := gjson.Get(ret, "data.result").String()
		lines := strings.Split(out, "\n")
		for _, v := range lines {
			fmt.Println(v)
		}
	}

	{
		//log
		req, _ := dhttp.NewRequest(fmt.Sprintf("http://%s/cmd/log", address), "POST")
		req.SetHeader("Access-Token", token)
		req.WriteJSON(struct {
			Name     string `json:"name"`
			PageNo   int    `json:"pageNo"`
			PageSize int    `json:"pageSize"`
		}{Name: "test", PageNo: 1, PageSize: 10})

		ret, err := req.ToString()
		t.Log(err, ret)
	}
}

func TestCmdHandler_Reg(t *testing.T) {
	str := `test {{name}}  {{na me}} 
{{ failed}} {{failed }} {{ failed }}
{{ failed} } { { failed}} { { failed} }
{{name}} 
{{{fds}} {{*}} {{9f}} {{_df}} {{+sf}} {{_}} {{f}} {{_9}} {{A6}}`
	// 以字母下划线开头，后接数字下划线和字母
	reg := regexp.MustCompile(`\{\{(_*[a-zA-Z]+[_a-zA-Z0-9]*)\}\}`)
	n := reg.FindAllString(str, -1)
	names := map[string]struct{}{}
	for _, name := range n {
		if _, ok := names[name]; !ok {
			names[name] = struct{}{}
		}
	}
	t.Log(n, names)
}
