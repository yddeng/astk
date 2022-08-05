package types

// 程序状态
type ProcStatus string

const (
	ProcStatusUnknown  ProcStatus = "unknown"
	ProcStatusStarting ProcStatus = "starting"
	ProcStatusRunning  ProcStatus = "running"
	ProcStatusStopping ProcStatus = "stopping"
	ProcStatusStopped  ProcStatus = "stopped"
	ProcStatusExited   ProcStatus = "exited"
)

type CmdType string

const (
	CmdTypeShell CmdType = "shell"
)

type GitType string

const (
	GitTypeGithub GitType = "github"
	GitTypeGitlab GitType = "gitlab"
)

type IncType string

const (
	IncTypeHttp  IncType = "http"
	IncTypeHttps IncType = "https"
	IncTypeTCP   IncType = "tcp"
)

type NotifyType string

const (
	// https://open.work.weixin.qq.com/help2/pc/14931?person_id=1&is_tencent=
	NotifyTypeWeixin NotifyType = "weixin"
	// 自定义类型
	NotifyTypeCallback NotifyType = "callback"
)
