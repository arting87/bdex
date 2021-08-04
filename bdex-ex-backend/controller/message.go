package controller

type BizMessage struct {
	Namespace string      `json:"namespace"`
	Action    string      `json:"action"`
	Data      interface{} `json:"data"` // data数据字段
}
