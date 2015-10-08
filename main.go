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

/*
import (
	"net/http"
	"log"
	"github.com/gorilla/websocket"
	"./client"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
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

		go client.New(101, conn).Run()
	})
	http.ListenAndServe(":8080", nil)
}
*/
/*
import (
	"./authentication"
	_ "./notifications"
	_ "fmt"
	"github.com/gorilla/websocket"
	_ "log"
	"net/http"
	_ "time"
	_ "./provider"
	"./router"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		id, err := authentication.AuthenticateRequest(r.FormValue("AccessToken"))
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		if id == -1 {
			http.Error(w, "", 401)
			return
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		conn.ReadMessage()
		router.AddClient <- router.Client{id, conn}
	})

	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			id, err := authentication.AuthenticateRequest(r.FormValue("AccessToken"))
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}

			if id == -1 {
				http.Error(w, "", 401)
				return
			}

			conn, err := upgrader.Upgrade(w, r, nil)

			for {
				_, b, err := conn.ReadMessage()
				if err != nil {
					http.Error(w, err.Error(), 500)
					return
				}

				err = conn.WriteMessage(websocket.TextMessage, b)
				if err != nil {
					http.Error(w, err.Error(), 500)
					return
				}
			}

			/*
			if err != nil {
				http.Error(w, "", 401)
			} else {
				conn, err := upgrader.Upgrade(w, r, nil)
				if err != nil {
					log.Fatal(err)
				}

				for {
					_, b, _ := conn.ReadMessage()
					fmt.Println(string(b))
				}
			}
*/

/*
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
*/
//})

//fmt.Println("listening on 8080")
//http.ListenAndServe(":8080", nil)
//}
