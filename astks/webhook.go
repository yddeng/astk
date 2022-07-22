package astks

type WebHook struct {
	Type     string `json:"type"`
	Key      string `json:"key"` // 由服务器创建
	Describe string `json:"describe"`
	Token    string `json:"token"`
}
