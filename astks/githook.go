package astks

import (
	"fmt"
	"github.com/go-playground/webhooks/github"
	"github.com/go-playground/webhooks/gitlab"
	"github.com/yddeng/astk/pkg/types"
)

type GitHook struct {
	ID      string        `json:"id"` // 服务创建
	Type    types.GitType `json:"type"`
	Name    string        `json:"name"`    // 仓库名
	Address string        `json:"address"` // 仓库地址
	Branch  string        `json:"branch"`  // 分支
	Token   string        `json:"token"`
	WebHook string        `json:"webHook"` // 地址
	Notify  Notify        `json:"notify"`

	Github *github.Webhook `json:"-"`
	Gitlab *gitlab.Webhook `json:"-"`
}

type GitHookMgr struct {
	Hooks map[string]*GitHook `json:"hooks"` // id -> GitHook
}

func (this *GitHook) makePushMessage(
	name, branch, user string,
	cmtIds, cmtAuthor, cmtMsgs []string, cmtCnt int,
) string {

	title := "[%s:%s] %s 推送了 %d 个提交"
	cmt := "  %s %s - %s"

	// 仅显示最近10个提交
	if len(cmtIds) > 10 {
		cmtIds = cmtIds[0:10]
	}

	ret := fmt.Sprintf(title, name, branch, user, cmtCnt)
	for i, id := range cmtIds {
		ret += "\n" + fmt.Sprintf(cmt, id[0:8], cmtAuthor[i], cmtMsgs[i])
	}
	if cmtCnt > len(cmtIds) {
		ret += "\n" + "..."
	}
	return ret
}

func (this *GitHook) makeMergeMessage(
	name, branch, user string,
	cmtIds, cmtAuthor, cmtMsgs []string,
) string {

	title := "[%s] %s 创建一个合并请求 %s->%s"
	cmt := "  %s %s - %s"

	ret := fmt.Sprintf(title, name, branch, user, len(cmtIds))
	for i, id := range cmtIds {
		ret += "\n" + fmt.Sprintf(cmt, id[0:8], cmtAuthor[i], cmtMsgs[i])
	}
	return ret
}
