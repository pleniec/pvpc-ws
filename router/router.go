package router

import (
	"github.com/gorilla/websocket"
)

var (
	clients map[int64]*websocket.Conn
)

func init() {
	clients = make(map[int64]*websocket.Conn)

	go func() {

	}()
}
