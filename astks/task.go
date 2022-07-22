package astks

import (
	"github.com/yddeng/timer"
	"time"
)

// 任务

type TaskType string

const (
	TaskTypeSecond TaskType = "second"
	TaskTypeDay    TaskType = "day"
	TaskTypeWeek   TaskType = "week"
	TaskTypeMonth  TaskType = "month"
)

type Task struct {
	Type    TaskType `json:"type"`    // 周期类型 second, day, week, month
	Begin   int64    `json:"begin"`   // 首次启动时间
	Loop    int64    `json:"loop"`    // 循环间隔，仅类型second使用，单位秒
	Times   int      `json:"times"`   // 执行次数
	DoTimes int      `json:"doTimes"` // 已经执行次数
	Next    int64    `json:"next"`    // 下一次执行时间

	timer    timer.Timer
	nextTime time.Time
}

var timerMgr = timer.NewHeapTimerMgr()

func CreateTask(task *Task, f func()) {
	if task.Times != 0 && task.DoTimes >= task.Times {
		task.Next = 0
		return
	}

	now := NowUnix()
	if task.Begin == 0 {
		task.Begin = now
	}

	task.Next = task.Begin
	task.nextTime = time.Unix(task.Next, 0)

	// 计算下次启动时间
	for now > task.Next {
		switch task.Type {
		case TaskTypeSecond:
			task.nextTime = task.nextTime.Add(time.Duration(task.Loop) * time.Second)
		case TaskTypeDay:
			task.nextTime = task.nextTime.AddDate(0, 0, 1)
		case TaskTypeWeek:
			task.nextTime = task.nextTime.AddDate(0, 0, 7)
		case TaskTypeMonth:
			task.nextTime = task.nextTime.AddDate(0, 1, 0)
		}
		task.Next = task.nextTime.Unix()
	}

	d := task.Next - now
	task.timer = timerMgr.OnceTimer(time.Duration(d)*time.Second, func() {
		taskQueue.Submit(func() {
			task.DoTimes++
			f()
			CreateTask(task, f)
		})
	})
}
