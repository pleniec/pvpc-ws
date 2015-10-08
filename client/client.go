package client

import (
	"encoding/json"
	_ "encoding/json"
	_ "fmt"
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	ID         int
	Connection *websocket.Conn
	InputChan  chan *Message
	OutputChan chan *Event
	ErrorChan  chan *Message
}

func New(id int, connection *websocket.Conn) *Client {
	return &Client{id, connection, make(chan *Message, 10), make(chan *Event, 10), make(chan *Message, 10)}
}

func (c *Client) Run() {
	//go c.readMessages()

	for {
		select {
		//case m := <- c.InputChan:
		//	fmt.Println(m)
		//	break
		case e := <-c.OutputChan:
			bytes, err := json.Marshal(e)
			if err != nil {
				log.Fatal(err)
			}

			c.Connection.WriteMessage(websocket.TextMessage, bytes)
			break
		}
	}
}
