package astks

import (
	"fmt"
	"github.com/go-playground/webhooks/github"
	"github.com/go-playground/webhooks/gitlab"
	token2 "github.com/yddeng/astk/pkg/token"
	"github.com/yddeng/astk/pkg/types"
	"log"
	"sort"
	"strings"
)

type githookHandler struct {
}

func (*githookHandler) genWebHook(key string) string {
	return fmt.Sprintf("http://%s:%d/githook/s/%s", config.Ip, config.WebConfig.Port, key)
}

func (*githookHandler) List(wait *WaitConn, user string, req struct {
	PageNo   int `json:"pageNo"`
	PageSize int `json:"pageSize"`
}) {
	//log.Printf("%s by(%s) %v\n", done.route, user, req)

	s := make([]*GitHook, 0, len(gitHookMgr.Hooks))
	for _, n := range gitHookMgr.Hooks {
		s = append(s, n)
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i].ID > s[j].ID
	})

	start, end := listRange(req.PageNo, req.PageSize, len(nodeMgr.Nodes))
	wait.SetResult("", pageData{
		PageNo:     req.PageNo,
		PageSize:   req.PageSize,
		TotalCount: len(s),
		Data:       s[start:end],
	})
	wait.Done()
}

func (this *githookHandler) Create(wait *WaitConn, user string, req struct {
	Type    types.GitType `json:"type"`
	Name    string        `json:"name"`
	Address string        `json:"address"` // 仓库地址
	Branch  string        `json:"branch"`  // 分支
	Token   string        `json:"token"`
	Notify  Notify        `json:"notify"`
}) {
	log.Printf("%s by(%s) %v\n", wait.route, user, req)
	defer func() { wait.Done() }()

	hook := &GitHook{
		Type:    req.Type,
		Name:    req.Name,
		Address: req.Address,
		Branch:  req.Branch,
		Token:   req.Token,
		Notify:  req.Notify,
	}

	id := token2.GenToken(22)
	for {
		if _, ok := gitHookMgr.Hooks[id]; !ok {
			break
		} else {
			id = token2.GenToken(22)
		}
	}

	hook.ID = id
	hook.WebHook = this.genWebHook(id)

	gitHookMgr.Hooks[id] = hook
	saveStore(snGitHookMgr)

}

func (*githookHandler) Delete(wait *WaitConn, user string, req struct {
	ID string `json:"id"`
}) {
	log.Printf("%s by(%s) %v\n", wait.route, user, req)
	defer func() { wait.Done() }()

	_, ok := gitHookMgr.Hooks[req.ID]
	if !ok {
		wait.SetResult("操作对像不存在", nil)
		return
	}

	delete(gitHookMgr.Hooks, req.ID)
	saveStore(snGitHookMgr)

}

var (
	githubEvents = []github.Event{github.PushEvent}
	gitlabEvents = []gitlab.Event{gitlab.PushEvents, gitlab.MergeRequestEvents}
)

func (this *githookHandler) Hook(wait *WaitConn) {
	log.Printf("%s \n", wait.route)

	defer func() { wait.Done() }()

	ctx := wait.Context()
	key := ctx.Param("key")

	hook, ok := gitHookMgr.Hooks[key]
	if !ok {
		return
	}

	switch hook.Type {
	case types.GitTypeGithub:
		if hook.Github == nil {
			var err error
			if hook.Github, err = github.New(github.Options.Secret("")); err != nil {
				log.Println(err)
				return
			}
		}

		payload, err := hook.Github.Parse(ctx.Request, githubEvents...)
		if err != nil {
			log.Println(err)
			if err == github.ErrEventNotFound {
				// ok event wasn;t one of the ones asked to be parsed
				return
			}
		}

		switch payload.(type) {
		case github.PushPayload:
			release := payload.(github.PushPayload)

			cmtIds := make([]string, 0, len(release.Commits))
			cmtAus := make([]string, 0, len(release.Commits))
			cmtMsgs := make([]string, 0, len(release.Commits))
			for _, cmt := range release.Commits {
				cmtIds = append(cmtIds, cmt.ID)
				cmtAus = append(cmtAus, cmt.Author.Name)
				cmtMsgs = append(cmtMsgs, cmt.Message)
			}

			msg := hook.makePushMessage(release.Repository.Name,
				strings.TrimPrefix(release.Ref, "refs/heads/"),
				release.Pusher.Name,
				cmtIds, cmtAus, cmtMsgs, len(cmtIds))
			log.Println(msg)

		case github.PullRequestPayload:
			pullRequest := payload.(github.PullRequestPayload)
			// Do whatever you want from here...
			fmt.Printf("%+v", pullRequest)
		}

	case types.GitTypeGitlab:
		if hook.Gitlab == nil {
			var err error
			if hook.Gitlab, err = gitlab.New(gitlab.Options.Secret("")); err != nil {
				log.Println(err)
				return
			}
		}

		payload, err := hook.Gitlab.Parse(ctx.Request, gitlabEvents...)
		if err != nil {
			log.Println(err)
			if err == github.ErrEventNotFound {
				// ok event wasn;t one of the ones asked to be parsed
				return
			}
		}

		switch payload.(type) {
		case gitlab.PushEventPayload:
			release := payload.(gitlab.PushEventPayload)

			cmtIds := make([]string, 0, len(release.Commits))
			cmtAus := make([]string, 0, len(release.Commits))
			cmtMsgs := make([]string, 0, len(release.Commits))
			for _, cmt := range release.Commits {
				cmtIds = append(cmtIds, cmt.ID)
				cmtAus = append(cmtAus, cmt.Author.Name)
				cmtMsgs = append(cmtMsgs, cmt.Message)
			}

			msg := hook.makePushMessage(release.Repository.Name,
				strings.TrimPrefix(release.Ref, "refs/heads/"),
				release.UserName,
				cmtIds, cmtAus, cmtMsgs, int(release.TotalCommitsCount))
			log.Println(msg)

		case gitlab.MergeRequestEventPayload:
			pullRequest := payload.(gitlab.MergeRequestEventPayload)
			// Do whatever you want from here...
			fmt.Printf("%+v", pullRequest)
		}
	default:

	}
}
