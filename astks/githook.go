package astks

import (
	"github.com/go-playground/webhooks/github"
	"github.com/go-playground/webhooks/gitlab"
	"github.com/yddeng/astk/pkg/types"
)

type GitHook struct {
	ID      string        `json:"id"` // 服务创建
	Type    types.GitType `json:"type"`
	Name    string        `json:"name"`    // 仓库名
	Address string        `json:"address"` // 仓库地址
	Token   string        `json:"token"`
	WebHook string        `json:"webHook"` // 地址
	Notify  Notify        `json:"notify"`

	Github *github.Webhook `json:"-"`
	Gitlab *gitlab.Webhook `json:"-"`
}

type GitHookMgr struct {
	Hooks map[string]*GitHook `json:"hooks"` // id -> GitHook
}
