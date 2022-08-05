package astks

import (
	"fmt"
	"log"
	"strings"
	"time"
)

const (
	alertCPU  = "cpu使用率超过 %.1f%%, 当前使用率 %.1f%%"
	alertMem  = "内存繁忙超过 %.1f%%, 当前使用率 %.1f%%"
	alertDisk = "硬盘使用率超过 %.1f%%, 当前使用率 %.1f%%"

	alertMessage = `告警内容：%s
告警对象：%s
触发时间：%s`
)

type Monitor struct {
	Cpu  int `json:"cpu"`  // cpu使用率, 默认90, 小于0表示不启用检测
	Mem  int `json:"mem"`  // mem使用率, 默认90, 小于0表示不启用检测
	Disk int `json:"disk"` // mem使用率, 默认90, 小于0表示不启用检测
	// 检测间隔， 默认10秒
	// 一段时间内，资源消耗都需要到达触发条件才报警
	Interval int64 `json:"interval"`
	// 下次报警时间间隔， 如果恢复了就重置
	// 为了防止持续报警轰炸
	AlertInterval int64 `json:"continuityInterval"`

	Opened bool   `json:"opened"`
	Notify Notify `json:"notify"` // 报警器
}

type MonitorState struct {
	TriggerTime int64 // 触发时间，第一次
	AlertTime   int64 // 报警时间
}

func (this *Monitor) trigger(cpu, mem, disk float64) (broken bool, info []string) {
	if this.Cpu > 0 && int(cpu) > this.Cpu {
		broken = true
		info = append(info, fmt.Sprintf(alertCPU, float64(this.Cpu), cpu))
	}
	if this.Mem > 0 && int(mem) > this.Mem {
		broken = true
		info = append(info, fmt.Sprintf(alertMem, float64(this.Mem), mem))
	}
	if this.Disk > 0 && int(disk) > this.Disk {
		broken = true
		info = append(info, fmt.Sprintf(alertDisk, float64(this.Disk), disk))
	}
	return
}

func (this *Monitor) Alert(state *MonitorState, cpu, mem, disk float64, name func() string) {
	broken, info := this.trigger(cpu, mem, disk)
	if !broken {
		state.TriggerTime = 0
		state.AlertTime = 0
	} else {
		now := time.Now()
		nowUnix := now.Unix()
		if state.TriggerTime == 0 {
			state.TriggerTime = nowUnix
		} else if nowUnix-state.TriggerTime >= this.Interval {
			// 持续时间内，已达到报警条件
			if nowUnix-state.AlertTime >= this.AlertInterval {
				//
				state.AlertTime = nowUnix
				if this.Notify.NotifyServer != "" {
					msg := fmt.Sprintf(alertMessage, strings.Join(info, ";"), name(), now.String())
					if err := this.Notify.Push(msg); err != nil {
						log.Println(err)
					}
				}
			}
		}
	}
}
