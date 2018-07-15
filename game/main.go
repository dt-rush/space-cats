package main

import (
	"fmt"
	"github.com/dt-rush/sameriver/engine"
	"github.com/dt-rush/space-cats/game/scene"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	fmt.Println("space cats, motherfucker")
	engine.RunGame(engine.GameInitSpec{
		WindowSpec: engine.WindowSpec{
			Title:      "space cats",
			Width:      400,
			Height:     300,
			Fullscreen: false},
		LoadingScene: &scene.LoadingScene{},
		FirstScene:   &scene.GameScene{},
	})
}
