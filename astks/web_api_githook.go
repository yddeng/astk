package astks

import (
	"fmt"
	"github.com/go-playground/webhooks/github"
	token2 "github.com/yddeng/astk/pkg/token"
	"log"
	"sort"
)

var (
	githubEvents = []github.Event{github.PushEvent, github.PullRequestEvent}
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
		return s[i].Key > s[j].Key
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
	Type    GitType `json:"type"`
	Name    string  `json:"name"`
	Address string  `json:"address"` // 仓库地址
	Token   string  `json:"token"`
	Notify  Notify  `json:"notify"`
}) {
	log.Printf("%s by(%s) %v\n", wait.route, user, req)

	hook := &GitHook{
		Type:    req.Type,
		Name:    req.Name,
		Address: req.Address,
		Token:   req.Token,
		Notify:  req.Notify,
	}

	hookKey := token2.GenToken(22)
	for {
		if _, ok := gitHookMgr.Hooks[hookKey]; !ok {
			break
		} else {
			hookKey = token2.GenToken(22)
		}
	}

	hook.Key = hookKey
	hook.WebHook = this.genWebHook(hookKey)

	gitHookMgr.Hooks[hookKey] = hook
	saveStore(snGitHookMgr)

	wait.Done()

}

func (this *githookHandler) Hook(wait *WaitConn) {
	log.Printf("%s \n", wait.route)

	defer func() { wait.Done() }()

	ctx := wait.Context()
	key := ctx.Param("key")

	hook, ok := gitHookMgr.Hooks[key]
	if !ok {
		return
	}

	fmt.Println(key, hook)

	switch hook.Type {
	case GitTypeGithub:
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
			// Do whatever you want from here...
			fmt.Printf("%+v", release)

		case github.PullRequestPayload:
			pullRequest := payload.(github.PullRequestPayload)
			// Do whatever you want from here...
			fmt.Printf("%+v", pullRequest)
		}

	case GitTypeGitlab:
	default:

	}
}
