package server

import (
	"log"
	"net/http"

	"github.com/SuperTikuwa/websocket-game/utility"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var jsonChannel = make(chan utility.Message)
var connections = make([]*websocket.Conn, 0)

func GetJson(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	connections = append(connections, c)
	go broadcast()

	for {
		var msg utility.Message
		err := c.ReadJSON(&msg)
		if err != nil {
			log.Fatal(err)
		}
		jsonChannel <- msg
	}
}

func broadcast() {
	for {
		json := <-jsonChannel
		for _, con := range connections {
			con.WriteJSON(json)
		}
	}
}
