package astks

import "log"

type MsgNotifyType string

const (
	MsgNotifyTypeWeixin MsgNotifyType = "weixin"
)

type MsgNotify interface {
	Name() string
	Type() MsgNotifyType
	Push(msg string) error
}

// https://open.work.weixin.qq.com/help2/pc/14931?person_id=1&is_tencent=
type WeixinNotify struct {
	url string
}

func (this *WeixinNotify) Name() string {
	return "test"
}

func (this *WeixinNotify) Type() MsgNotifyType {
	return MsgNotifyTypeWeixin
}

func (this *WeixinNotify) Push(msg string) error {
	log.Println("weixinRobot push", msg)
	return nil
}
