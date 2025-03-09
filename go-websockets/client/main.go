package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/gorilla/websocket"
)

type Message struct {
	MessageType int
	Data        []byte
}

func main() {

	// connecting to the remote ws
	u := url.URL{Scheme: "ws", Host: "localhost:3000", Path: "/ws"}
	fmt.Printf("Connecting to %s\n", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

	if err != nil {
		log.Fatal("dial:", err)
	}

	defer conn.Close()

	// channels for managing the messages
	send := make(chan Message)
	done := make(chan struct{})

	// Goroutine for reading messages
	go func() {
		defer close(done)
		for {
			_, messages, err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}

			fmt.Printf("Received: %s\n", messages)
		}
	}()

	// Goroutine for sending messages
	go func() {
		for {
			select {
			case msg := <-send:
				// write the message to the websocket connection
				err := conn.WriteMessage(msg.MessageType, msg.Data)
				if err != nil {
					log.Println("write err", err)
					return
				}

			case <-done:
				return
			}
		}
	}()

	// read input from the terminal and send it to the websocket server

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Type something...")
	for scanner.Scan() {
		text := scanner.Text()
		// send the message to the server
		send <- Message{websocket.TextMessage, []byte(text)}
	}

	if err := scanner.Err(); err != nil {
		log.Println("Scanner error	", err)
	}

}
