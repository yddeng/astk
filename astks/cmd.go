package astks

import "regexp"

type Cmd struct {
	ID       int                 `json:"id"`
	Name     string              `json:"name"`
	Dir      string              `json:"dir"`
	Context  string              `json:"context"`
	Args     map[string]string   `json:"args"`
	User     string              `json:"user"`
	UpdateAt int64               `json:"update_at"`
	CreateAt int64               `json:"create_at"`
	CallNo   int                 `json:"call_no"`
	doing    map[string]struct{} // 节点正在执行
}

type CmdMgr struct {
	Success int               `json:"success"`
	Failed  int               `json:"failed"`
	GenID   int               `json:"gen_id"`
	CmdMap  map[int]*Cmd      `json:"cmd_map"`
	CmdLogs map[int][]*CmdLog `json:"cmd_logs"`
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
	CreateAt int64  `json:"create_at"` // 执行时间
	User     string `json:"user"`      // 执行用户
	Dir      string `json:"dir"`       // 执行目录
	Node     string `json:"node"`      // 执行的节点
	Timeout  int    `json:"timeout"`   // 执行超时时间
	Context  string `json:"context"`   // 执行内容
	ResultAt int64  `json:"result_at"` // 执行结果时间
	Result   string `json:"result"`    // 执行结果
}
