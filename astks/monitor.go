package astks

import "log"

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

	Notifys []string `json:"notifys"` // 报警器
}

type MonitorState struct {
	TriggerTime int64 // 触发时间，第一次
	AlertTime   int64 // 报警时间
}

func (this *Monitor) Trigger(cpu, mem, disk float64) bool {
	if this.Cpu > 0 && int(cpu) > this.Cpu {
		return true
	}
	if this.Mem > 0 && int(mem) > this.Mem {
		return true
	}
	if this.Disk > 0 && int(disk) > this.Disk {
		return true
	}
	return false
}

func (this *Monitor) Alert(state *MonitorState, trigger bool) {
	if !trigger {
		state.TriggerTime = 0
		state.AlertTime = 0
	} else {
		now := NowUnix()
		if state.TriggerTime == 0 {
			state.TriggerTime = now
		} else if now-state.TriggerTime > this.Interval {
			// 持续时间内，已达到报警条件
			if now-state.AlertTime > this.AlertInterval {
				//
				log.Println("alert ")
				state.AlertTime = now
			}
		}
	}
}
