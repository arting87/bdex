package ws

import (
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type WSConnection struct {
	mutex             sync.Mutex
	connAddr          string
	connId            uint64
	wsSocket          *websocket.Conn
	inChan            chan BizMessage
	outChan           chan BizMessage
	closeChan         chan byte
	isClosed          bool
	lastHeartbeatTime time.Time
}

// 读websocket
func (wsConnection *WSConnection) readLoop() {
	var (
		bm  BizMessage
		err error
		p   []byte
	)
	log.Infow("开始读循环", "connAddr", wsConnection.connAddr)
	for {
		if _, p, err = wsConnection.wsSocket.ReadMessage(); err != nil {
			log.Errorw("读取信息失败", "error", err)
			goto ERR
		}
		//log.Infow("读到的信息", "connAddr", wsConnection.connAddr, "message", p)

		err = json.Unmarshal(p, &bm)
		if err != nil {
			log.Infow("json解析请求数据失败", "error", err)
			goto ERR
		}
		//log.Infow("json解析后的信息", "connAddr", wsConnection.connAddr, bm.Namespace+"_"+bm.Action, bm.Data)

		select {
		case wsConnection.inChan <- bm:
		case <-wsConnection.closeChan:
			goto CLOSED
		}
	}

ERR:
	wsConnection.Close()
CLOSED:
	wsConnection.Close()
}

// 写websocket
func (wsConnection *WSConnection) writeLoop() {
	var (
		message BizMessage
		err     error
	)
	for {
		select {
		case message = <-wsConnection.outChan:
			if err = wsConnection.wsSocket.WriteJSON(message); err != nil {
				goto ERR
			}
		case <-wsConnection.closeChan:
			goto CLOSED
		}
	}
ERR:
	wsConnection.Close()
CLOSED:
}

/**
以下是API
*/

func InitWSConnection(connId uint64, connAddr string, wsSocket *websocket.Conn) (wsConnection *WSConnection) {
	wsConnection = &WSConnection{
		wsSocket:          wsSocket,
		connAddr:          connAddr,
		connId:            connId,
		inChan:            make(chan BizMessage, WsInChannelSize),
		outChan:           make(chan BizMessage, WsOutChannelSize),
		closeChan:         make(chan byte),
		lastHeartbeatTime: time.Now(),
	}
	log.Infow("开始循环读写", "connAddr", connAddr)
	go wsConnection.readLoop()
	go wsConnection.writeLoop()

	return
}

// 发送消息
func (wsConnection *WSConnection) SendMessage(message BizMessage) (err error) {
	select {
	case wsConnection.outChan <- message:
		if message.Namespace != api {
			if message.Action != get {
				log.Infow("发出信息", "connAddr", wsConnection.connAddr, "message", message)
			}
		}
		SendMessageTotal_INCR()
	case <-wsConnection.closeChan:
		err = errors.New("ERR_CONNECTION_LOSS")
	default: // 写操作不会阻塞, 因为channel已经预留给websocket一定的缓冲空间
		err = errors.New("ERR_SEND_MESSAGE_FULL")
		SendMessageFail_INCR()
	}
	return
}

// 读取消息
func (wsConnection *WSConnection) ReadMessage() (message BizMessage, err error) {
	select {
	case message = <-wsConnection.inChan:
	case <-wsConnection.closeChan:
		err = errors.New("ERR_CONNECTION_LOSS")
	}
	return
}

// 关闭连接
func (wsConnection *WSConnection) Close() {
	wsConnection.wsSocket.Close()

	wsConnection.mutex.Lock()
	defer wsConnection.mutex.Unlock()

	if !wsConnection.isClosed {
		wsConnection.isClosed = true
		close(wsConnection.closeChan)
	}
}

// 检查心跳（不需要太频繁）
func (wsConnection *WSConnection) IsAlive() bool {
	var (
		now = time.Now()
	)

	wsConnection.mutex.Lock()
	defer wsConnection.mutex.Unlock()

	// 连接已关闭 或者 太久没有心跳
	if wsConnection.isClosed || now.Sub(wsConnection.lastHeartbeatTime) > time.Duration(WsHeartbeatInterval)*time.Second {
		return false
	}
	return true
}

// 更新心跳
func (wsConnection *WSConnection) KeepAlive() {
	var (
		now = time.Now()
	)

	wsConnection.mutex.Lock()
	defer wsConnection.mutex.Unlock()

	wsConnection.lastHeartbeatTime = now
}
