package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var connections = make(map[int]*websocket.Conn)
var count = 0

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc("/", pingPong)
	http.ListenAndServe(":12345", nil)
}

func pingPong(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	connections[count] = conn
	count++
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, connection := range connections {
			_ = connection.WriteMessage(msgType, []byte(string(msg)))
		}
	}
}
