package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

func main() {
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer c.Close()

	// Graceful shutdown on Ctrl+C
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Send message
	err = c.WriteMessage(websocket.TextMessage, []byte("Hello from client!"))
	if err != nil {
		log.Fatal("Write error:", err)
	}

	// Wait for reply
	_, msg, err := c.ReadMessage()
	if err != nil {
		log.Fatal("Read error:", err)
	}
	fmt.Println("Received from server:", string(msg))
}
