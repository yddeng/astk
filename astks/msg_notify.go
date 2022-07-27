package astks

import "log"

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
		this.pushWeixinMessage(msg)
	case MsgNotifyTypeCallback:
		this.pushCallbackMessage(msg)
	}

	return nil
}

func (this *Notify) pushWeixinMessage(msg string) {
	log.Println("weixin push", msg)
}

func (this *Notify) pushCallbackMessage(msg string) {
	log.Println("callback push", msg)

}
