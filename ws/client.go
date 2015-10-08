package ws

import (
	"fmt"
	"log"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID int
	Connection *websocket.Conn
	Router *Router
}

func NewClient(id int, connection *websocket.Conn, router *Router) *Client {
	return &Client{id, connection, router}
}

func (c *Client) Run() {
	go c.readMessages()

	select {
	}
}

func (c *Client) readMessages() {
	for {
		mt, b, err := c.Connection.ReadMessage()
		if err != nil {
			if err.(*websocket.CloseError) != nil {
				log.Fatal("CLOSED")
			} else {
				log.Fatal(err)
			}
		}

		fmt.Println(mt)
		fmt.Println(string(b))
	}
}

/*
package client

import (
	"fmt"
	"log"
	"github.com/gorilla/websocket"
	"../router"
)

type Client struct {
	ID int
	Connection *websocket.Conn
}

func New(id int, conn *websocket.Conn) (*Client) {
	return &Client{id, conn}
}

func (c *Client) Run() {
	router.AddClient < c

	go c.readMessages()

	select {
	}
}

func (c *Client) readMessages() {
	for {
		mt, b, err := c.Connection.ReadMessage()
		if err != nil {
			if err.(*websocket.CloseError) != nil {
				log.Fatal("CLOSED")
			} else {
				log.Fatal(err)
			}
		}

		fmt.Println(mt)
		fmt.Println(string(b))
	}
}

*/