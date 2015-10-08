package router

import (
	"github.com/gorilla/websocket"
)

var (
	clients map[int64]*websocket.Conn
	AddClient chan struct{int; *websocket.Conn}
)

func init() {
	clients = make(map[int64]*websocket.Conn)
	AddClient = make(chan struct{int; *websocket.Conn}, 10)

	go func() {

	}()
}
