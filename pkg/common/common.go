package common

import "time"

// 隐藏目录，存放临时配置文件
const Dir = ".astk"

// 程序状态
const (
	StateUnknown  = "unknown"
	StateStarting = "starting"
	StateRunning  = "running"
	StateStopping = "stopping"
	StateStopped  = "stopped"
	StateExited   = "exited"
)

const HeartbeatTimeout = 10 * time.Second
