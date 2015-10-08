package ws

import (
	"fmt"
	_ "github.com/gorilla/websocket"
)

type Router struct {
	Clients map[int]*Client
	AddClient chan *Client
	RemoveClient chan *Client
}

func NewRouter() *Router {
	return &Router{make(map[int]*Client), make(chan *Client), make(chan *Client)}
}

func (r *Router) Run() {
	for {
		select {
		case c := <- r.AddClient:
			fmt.Println("added client (ID: %d)", c.ID)
			r.Clients[c.ID] = c
			break
		case c := <- r.RemoveClient:
			fmt.Println("removed client (ID: %d)", c.ID)
			delete(r.Clients, c.ID)
			break
		}
	}
}
