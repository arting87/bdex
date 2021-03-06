package ws

import (
	"encoding/json"
	"sync/atomic"
)

type Stats struct {
	// 反馈在线长连接的数量
	OnlineConnections int64 `json:"onlineConnections"`

	// 反馈客户端的推送压力
	SendMessageTotal int64 `json:"sendMessageTotal"`
	SendMessageFail  int64 `json:"sendMessageFail"`
}

var (
	G_stats *Stats
)

func InitStats() (err error) {
	G_stats = &Stats{}
	return
}

func OnlineConnections_INCR() {
	atomic.AddInt64(&G_stats.OnlineConnections, 1)
}

func OnlineConnections_DESC() {
	atomic.AddInt64(&G_stats.OnlineConnections, -1)
}

func SendMessageFail_INCR() {
	atomic.AddInt64(&G_stats.SendMessageFail, 1)
}

func SendMessageTotal_INCR() {
	atomic.AddInt64(&G_stats.SendMessageTotal, 1)
}

func (stats *Stats) Dump() (data []byte, err error) {
	return json.Marshal(G_stats)
}
