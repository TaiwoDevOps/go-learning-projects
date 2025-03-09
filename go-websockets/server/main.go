package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Received : %s\n", message)

		res := fmt.Sprintf("%s, at: %s\n", string(message), time.Now().String())

		if err := conn.WriteMessage(messageType, []byte(res)); err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main() {

	http.HandleFunc("/ws", handler)
	fmt.Println("Starting server on port: 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println(err)
	}

}
