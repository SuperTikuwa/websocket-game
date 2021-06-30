package main

import (
	"net/http"

	"github.com/SuperTikuwa/websocket-game/server/server"
)

func main() {
	http.HandleFunc("/", server.GetJson)
	http.ListenAndServe(":8080", nil)
}
