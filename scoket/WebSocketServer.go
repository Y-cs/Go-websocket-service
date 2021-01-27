package scoket

import (
	_ "GoTest/redis"
	"GoTest/service"
	_ "GoTest/service"
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	//websocket 升级器
	upgrade = websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	//动作处理
	action = service.Action{}
)

func WsHandler(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("hello"))
	var (
		wsConn *websocket.Conn
		err    error
		conn   *Connection
		data   []byte
	)
	// 完成ws协议的握手操作
	// Upgrade:websocket
	if wsConn, err = upgrade.Upgrade(w, r, nil); err != nil {
		return
	}

	if conn, err = InitConnection(wsConn); err != nil {
		goto ERR
	}

	//// 启动线程，不断发消息
	//go func() {
	//	var (
	//		err error
	//	)
	//	for {
	//		if err = conn.WriteMessage([]byte("heartbeat")); err != nil {
	//			return
	//		}
	//		time.Sleep(30 * time.Second)
	//	}
	//}()

	for {
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		//处理动作
		action.ActionHandle(data)
	}

ERR:
	conn.Close()

}

