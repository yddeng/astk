package astks

import "log"

type MsgRobotType string

const (
	MsgRobotTypeWeixin MsgRobotType = "weixin"
)

type MsgRobot interface {
	Type() MsgRobotType
	Push(msg string) error
}

// https://open.work.weixin.qq.com/help2/pc/14931?person_id=1&is_tencent=
type WeixinRobot struct {
	url string
}

func (this *WeixinRobot) Type() MsgRobotType {
	return MsgRobotTypeWeixin
}

func (this *WeixinRobot) Push(msg string) error {
	log.Println("weixinRobot push", msg)
	return nil
}
