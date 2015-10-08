package main

import (
	"./client"
	"./router"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func main() {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal(err)
		}

		c := client.New(101, conn)
		router.AddClientChan <- c
		go c.Run()
	})

	http.ListenAndServe(":8080", nil)
}
