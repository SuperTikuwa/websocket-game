package gamescene

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/SuperTikuwa/websocket-game/utility"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/supertikuwa/ebiten-scene/scene"
)

var gopher *ebiten.Image
var Player character

var Tick int

type character struct {
	Image *ebiten.Image
	Pos   utility.Position
}

func init() {
	file, err := os.Open("./gamescene/images/1000.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	gopher = ebiten.NewImageFromImage(img)

	Player = character{Image: gopher, Pos: utility.Position{X: 0, Y: 0}}
}

type GameScene struct {
}

func (g *GameScene) Update(sceneManager *scene.SceneManager) error {
	Tick++
	fmt.Println(Player.Image.Bounds().Dx())
	if Player.Pos.X >= 0 && Player.Pos.X <= 720-Player.Image.Bounds().Dx()/10 {
		if ebiten.IsKeyPressed(ebiten.KeyD) {
			// Player.Pos.X += inpututil.KeyPressDuration(ebiten.KeyD)
			Player.Pos.X += 3
		}
		if ebiten.IsKeyPressed(ebiten.KeyA) {
			// Player.Pos.X -= inpututil.KeyPressDuration(ebiten.KeyA)
			Player.Pos.X -= 3
		}
	} else {
		if Player.Pos.X < 0 {
			Player.Pos.X = 0
		} else if Player.Pos.X > 720-Player.Image.Bounds().Dx()/10 {
			Player.Pos.X = 720 - Player.Image.Bounds().Dx()/10
		}
	}

	if Player.Pos.Y >= 0 && Player.Pos.Y <= 480-Player.Image.Bounds().Dy()/10 {
		if ebiten.IsKeyPressed(ebiten.KeyW) {
			// Player.Pos.X -= inpututil.KeyPressDuration(ebiten.KeyA)
			Player.Pos.Y -= 3
		}
		if ebiten.IsKeyPressed(ebiten.KeyS) {
			// Player.Pos.X -= inpututil.KeyPressDuration(ebiten.KeyA)
			Player.Pos.Y += 3
		}
	} else {
		if Player.Pos.Y < 0 {
			Player.Pos.Y = 0
		} else if Player.Pos.Y > 480-Player.Image.Bounds().Dy()/10 {
			Player.Pos.Y = 480 - Player.Image.Bounds().Dy()/10
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		os.Exit(0)
	}

	return nil
}

func (g *GameScene) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.1, 0.1)
	op.GeoM.Translate(float64(Player.Pos.X), float64(Player.Pos.Y))
	screen.DrawImage(gopher, op)
}

func NewGameScene() *GameScene {
	return &GameScene{}
}
