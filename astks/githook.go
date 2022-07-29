package astks

import (
	"github.com/go-playground/webhooks/github"
	"github.com/go-playground/webhooks/gitlab"
)

type GitType string

const (
	GitTypeGithub GitType = "github"
	GitTypeGitlab GitType = "gitlab"
)

type GitHook struct {
	Type    GitType `json:"type"`
	Name    string  `json:"name"`    // 仓库名
	Address string  `json:"address"` // 仓库地址
	Token   string  `json:"token"`
	WebHook string  `json:"webHook"` // 地址
	Key     string  `json:"key"`     // 服务创建
	Notify  Notify  `json:"notify"`

	Github *github.Webhook `json:"-"`
	Gitlab *gitlab.Webhook `json:"-"`
}

type GitHookMgr struct {
	Hooks map[string]*GitHook `json:"hooks"` // key -> GitHook
}
