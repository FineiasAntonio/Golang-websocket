package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
}

func handler(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	log.Println("Lisening", conn.LocalAddr().Network())

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
		}

		conn.WriteMessage(messageType, message)
	}
}

func main() {

	http.HandleFunc("/ws", handler)
	http.ListenAndServe(":8080", nil)

}
