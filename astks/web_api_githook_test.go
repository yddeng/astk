package astks

import (
	"fmt"
	"github.com/go-playground/webhooks/github"
	"github.com/tidwall/gjson"
	"github.com/yddeng/dnet/dhttp"
	"net/http"
	"os"
	"testing"
)

func TestGithookHandler_Create(t *testing.T) {
	startWebListener(t)

	ret := authLogin(t, "admin", "123456")
	t.Log(ret, gjson.Get(ret, "data.token").String())

	req2, _ := dhttp.NewRequest(fmt.Sprintf("http://%s/githook/create", address), "POST")
	req2.SetHeader("Access-Token", gjson.Get(ret, "data.token").String())
	req2, _ = req2.WriteJSON(struct {
		Type    GitType `json:"type"`
		Name    string  `json:"name"`
		Address string  `json:"address"` // 仓库地址
		Token   string  `json:"token"`
		Notify  Notify  `json:"notify"`
	}{Type: GitTypeGithub,
		Name:    "test",
		Address: "https://github.com/yddeng/webhook",
		Token:   "123456",
		Notify: Notify{
			NotifyType:   MsgNotifyTypeCallback,
			NotifyServer: "http://127.0.0.1:24563/hook",
		}})

	ret, err := req2.ToString()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret)

	{
		req2, _ := dhttp.NewRequest(fmt.Sprintf("http://%s/githook/list", address), "POST")
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
}

func TestGithookHandler_List(t *testing.T) {
	startWebListener(t)

	ret := authLogin(t, "admin", "123456")
	t.Log(ret, gjson.Get(ret, "data.token").String())
	req2, _ := dhttp.NewRequest(fmt.Sprintf("http://%s/githook/list", address), "POST")
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

func TestGithookHandler_Hook(t *testing.T) {
	tests := []struct {
		name     string
		event    github.Event
		typ      interface{}
		filename string
		headers  http.Header
	}{
		{
			name:     "PullRequestEvent",
			event:    github.PullRequestEvent,
			typ:      github.PullRequestPayload{},
			filename: "../testdata/github/pull-request.json",
			headers: http.Header{
				"X-Github-Event":  []string{"pull_request"},
				"X-Hub-Signature": []string{"sha1=123456"},
			},
		},
		{
			name:     "PushEvent",
			event:    github.PushEvent,
			typ:      github.PushPayload{},
			filename: "../testdata/github/push.json",
			headers: http.Header{
				"X-Github-Event":  []string{"push"},
				"X-Hub-Signature": []string{"sha1=123456"},
			},
		},
	}

	url := "http://127.0.0.1:40156/githook/s/b9M0gDDBrxBI"

	for _, tt := range tests {
		tc := tt
		payload, _ := os.Open(tc.filename)
		req, err := http.NewRequest(http.MethodPost, url, payload)
		req.Header = tc.headers
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		fmt.Println(err)
		fmt.Println(resp.StatusCode)

	}
}
