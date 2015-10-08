package router

import (
	"fmt"
	"github.com/gorilla/websocket"
	"../client"
)

var (
	clients map[int64]*client.Client
	AddClient chan *client.Client
	RemoveClient chan *client.Client
)

func init() {
	clients = make(map[int64]*websocket.Conn)
	AddClient = make(chan client.Client, 10)
	RemoveClient = make(chan client.Client, 10)

	go func() {
		for {
			select {
			case c := <- AddClient:
				fmt.Printf("new client (ID: %d) connected\n", c.ID)
				clients[c.ID] = c
				break
			case c := <- RemoveClient:
				delete(clients, c.ID)
				break
			}
		}
	}()
}
