package main

import (
	"log"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	_ "gopkg.in/redis.v3"
	"time"
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

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal(err)
		}


		conn.SetPongHandler(func (ad string) error {
			conn.SetReadDeadline(time.Now().Add(5 * time.Second))
			fmt.Println("PONG")
			return nil
		})
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))

		go func() {
			for {
				_, b, err := conn.ReadMessage()
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("CO KURWA?")
				fmt.Println(string(b))
			}
		}()

		go func() {
			for {
				time.Sleep(time.Second)
				conn.SetWriteDeadline(time.Now().Add(1 * time.Second))
				conn.WriteMessage(websocket.PingMessage, []byte{})
			}
		}()
	})
	http.ListenAndServe(":8080", nil)
}
