package ws

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/websocket"
)

// 	WebSocket服务端
type WSServer struct {
	server *http.Server
}

var (
	G_wsServer *WSServer

	wsUpgrader = websocket.Upgrader{
		// 允许所有CORS跨域请求
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func handleConnect(resp http.ResponseWriter, req *http.Request) {
	log.Infow("New Connection", "Info:", req.RemoteAddr)
	var (
		err      error
		wsSocket *websocket.Conn
		connId   uint64
		wsConn   *WSConnection
	)

	// WebSocket握手
	if wsSocket, err = wsUpgrader.Upgrade(resp, req, nil); err != nil {
		return
	}

	// 连接唯一标识
	//connId = atomic.AddUint64(&G_wsServer.curConnId, 1)

	// 初始化WebSocket的读写协程
	wsConn = InitWSConnection(connId, req.RemoteAddr, wsSocket)

	// 开始处理websocket消息
	wsConn.WSHandle()
}

func InitWSServer() {
	// 统计
	var err error
	if err = InitStats(); err != nil {
		log.Errorw("init stats", "error", err)
		return
	}
	//var (
	//	mux      *http.ServeMux
	//	server   *http.Server
	//	listener net.Listener
	//)

	// 路由
	//mux = http.NewServeMux()
	//mux.HandleFunc("/ws", handleConnect)

	// HTTP服务
	//server = &http.Server{
	//	ReadTimeout:  time.Duration(WsReadTimeout) * time.Millisecond,
	//	WriteTimeout: time.Duration(WsWriteTimeout) * time.Millisecond,
	//	Handler:      mux,
	//}

	//fmt.Println(server.Addr)
	//
	//// 监听端口
	//if listener, err = net.Listen("tcp", ":"+strconv.Itoa(WsPort)); err != nil {
	//	log.Errorw("net listen", "error", err)
	//	return
	//}
	//fmt.Println(listener.Addr())
	//// 赋值全局变量
	//G_wsServer = &WSServer{
	//	server:    server,
	//	curConnId: uint64(time.Now().Unix()),
	//}
	//fmt.Println("拉起服务")
	//// 拉起服务
	//server.Serve(listener)
	http.HandleFunc("/ws", handleConnect)

	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "hello http")
	})

	fmt.Println("拉起服务在0.0.0.0:" + WsPort)
	err = http.ListenAndServe("0.0.0.0:"+WsPort, nil)
	fmt.Println("拉起服务错误：" + err.Error())
}
