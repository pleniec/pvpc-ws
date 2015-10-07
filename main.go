package main

import (
	"log"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type Message struct {
	Text string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal(err)
		}

		for {
			m := &Message{}
			err = conn.ReadJSON(m)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(*m)
		}

		conn.WriteMessage(websocket.TextMessage, []byte("sracz"))
	})
	http.ListenAndServe(":8080", nil)
}
