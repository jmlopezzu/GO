// main.go

package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	// Your implementation may vary, but you need to keep track of connected clients
	// and broadcast messages to all connected clients
	// See: https://github.com/gorilla/websocket/blob/master/examples/chat/main.go
}

func main() {
	// Configure the WebSocket route
	http.HandleFunc("/ws", handleConnections)

	// Start listening for incoming chat messages
	go handleMessages()

	// Start the server on localhost port 3000 and log any errors
	fmt.Println("WebSocket server started on :3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
