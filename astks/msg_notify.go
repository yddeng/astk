package astks

import (
	"fmt"
	"github.com/yddeng/dnet/dhttp"
)

type MsgNotifyType string

const (
	// https://open.work.weixin.qq.com/help2/pc/14931?person_id=1&is_tencent=
	MsgNotifyTypeWeixin MsgNotifyType = "weixin"
	// 自定义类型
	MsgNotifyTypeCallback MsgNotifyType = "callback"
)

type Notify struct {
	NotifyType   MsgNotifyType `json:"notifyType"`
	NotifyServer string        `json:"notifyServer"`
}

func (this *Notify) Type() MsgNotifyType {
	return this.NotifyType
}

func (this *Notify) Push(msg string) error {
	switch this.NotifyType {
	case MsgNotifyTypeWeixin:
		return this.pushWeixinMessage(msg)
	case MsgNotifyTypeCallback:
		return this.pushCallbackMessage(msg)
	default:
		return fmt.Errorf("push notify type %s invaild", this.NotifyType)
	}
}

type NotifyMsg struct {
	Message string `json:"message"`
}

func (this *Notify) pushWeixinMessage(msg string) error {
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
