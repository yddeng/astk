package astks

import (
	"fmt"
	"github.com/yddeng/astk/pkg/types"
	"github.com/yddeng/dnet/dhttp"
)

type Notify struct {
	NotifyType   types.NotifyType `json:"notifyType"`
	NotifyServer string           `json:"notifyServer"`
}

func (this *Notify) Type() types.NotifyType {
	return this.NotifyType
}

func (this *Notify) Push(msg string) error {
	switch this.NotifyType {
	case types.NotifyTypeWeixin:
		return this.pushWeixinMessage(msg)
	case types.NotifyTypeCallback:
		return this.pushCallbackMessage(msg)
	default:
		return fmt.Errorf("push notify type %s invaild", this.NotifyType)
	}
}

type NotifyMsg struct {
	Message string `json:"message"`
}

func (this *Notify) pushWeixinMessage(msg string) error {

	type WeixinMessage struct {
		Msgtype string `json:"msgtype"`
		Text    struct {
			Content       string   `json:"content"`
			MentionedList []string `json:"mentioned_list"`
		} `json:"text"`
	}

	data := WeixinMessage{
		Msgtype: "text",
		Text: struct {
			Content       string   `json:"content"`
			MentionedList []string `json:"mentioned_list"`
		}{
			Content: msg,
		},
	}

	req, err := dhttp.PostJson(this.NotifyServer, data)
	if err != nil {
		return err
	}

	resp, err := req.Do()
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("http post %s statecode %d", this.NotifyServer, resp.StatusCode)
	}
	return nil
}

func (this *Notify) pushCallbackMessage(msg string) error {
	req, err := dhttp.PostJson(this.NotifyServer, NotifyMsg{Message: msg})
	if err != nil {
		return err
	}

	resp, err := req.Do()
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("http post %s statecode %d", this.NotifyServer, resp.StatusCode)
	}
	return nil
}
