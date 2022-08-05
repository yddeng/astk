package astks

import (
	"github.com/yddeng/astk/pkg/types"
	"regexp"
)

type Cmd struct {
	ID       int                 `json:"id"`
	Type     types.CmdType       `json:"type"`
	Name     string              `json:"name"`
	Dir      string              `json:"dir"`
	Context  string              `json:"context"`
	Args     map[string]string   `json:"args"`
	User     string              `json:"user"`
	UpdateAt int64               `json:"updateAt"`
	CreateAt int64               `json:"createAt"`
	ExecAt   int64               `json:"execAt"`
	CallNo   int                 `json:"callNo"`
	doing    map[string]struct{} // 节点正在执行
}

type CmdMgr struct {
	Success  int               `json:"success"`
	Failed   int               `json:"failed"`
	GenID    int               `json:"genId"`
	CmdMap   map[int]*Cmd      `json:"cmdMap"`
	CmdLogs  map[int][]*CmdLog `json:"cmdLogs"`
	CmdTasks map[int]*CmdTask  `json:"cmdTasks"`
}

// 以字母下划线开头，后接数字下划线和字母
func cmdContextReg(str string) map[string]struct{} {
	reg := regexp.MustCompile(`\{\{(_*[a-zA-Z]+[_a-zA-Z0-9]*)\}\}`)
	n := reg.FindAllString(str, -1)
	names := map[string]struct{}{}
	for _, name := range n {
		if _, ok := names[name]; !ok {
			names[name] = struct{}{}
		}
	}
	return names
}

var cmdLogCapacity int = 10

type CmdLog struct {
	ID       int    `json:"id"`
	CreateAt int64  `json:"createAt"` // 执行时间
	User     string `json:"user"`     // 执行用户
	Dir      string `json:"dir"`      // 执行目录
	Node     string `json:"node"`     // 执行的节点
	Timeout  int    `json:"timeout"`  // 执行超时时间
	Context  string `json:"context"`  // 执行内容
	ResultAt int64  `json:"resultAt"` // 执行结果时间
	Result   string `json:"result"`   // 执行结果
}

type CmdTask struct {
	Task
	TaskID  int               `json:"taskId"`
	ID      int               `json:"id"`
	Dir     string            `json:"dir"`
	Args    map[string]string `json:"args"`
	Node    string            `json:"node"`
	Timeout int               `json:"timeout"`
}

func (this *CmdTask) Do() {

}
