package main

import (
	"log"

	"github.com/SuperTikuwa/websocket-game/client/gamescene"
	"github.com/SuperTikuwa/websocket-game/utility"
	"github.com/gorilla/websocket"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/supertikuwa/ebiten-scene/scene"
)

type Game struct {
	SceneManager *scene.SceneManager
}

func (g *Game) Update() error {
	if g.SceneManager == nil {
		g.SceneManager = &scene.SceneManager{}
		g.SceneManager.GoTo(gamescene.NewGameScene())
		return nil
	}

	g.SceneManager.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.SceneManager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 720, 480
}

var url = "ws://localhost:8080/"

func main() {
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal(err)
	}

	go receiver(c)
	go request(c)

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

}

func receiver(c *websocket.Conn) {

	for {
		var msg utility.Message
		err := c.ReadJSON(&msg)
		if err != nil {
			log.Fatal(err)
		}

		if gamescene.Tick == int(ebiten.CurrentTPS()) {
			sync(msg)
		}

		// log.Println(msg)
	}
}

func sync(msg utility.Message) {
	gamescene.Player.Pos = msg.Pos
}

func request(c *websocket.Conn) {
	for {
		msg := utility.Message{ID: 1, Pos: gamescene.Player.Pos}
		c.WriteJSON(msg)
	}
}
