package server

import (
	"log"
	"net/http"

	"github.com/SuperTikuwa/websocket-game/utility"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var msgChannel = make(chan utility.Message)

func GetJson(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	for {
		var msg utility.Message
		c.ReadJSON(&msg)
	}
}

func broadcast() {

}
