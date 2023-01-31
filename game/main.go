package main

import (
	"fmt"
	"github.com/dt-rush/sameriver/v2"
	"github.com/dt-rush/space-cats/game/scene"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	fmt.Println("space cats, motherfucker")
	sameriver.RunGame(sameriver.GameInitSpec{
		WindowSpec: sameriver.WindowSpec{
			Title:      "space cats",
			Width:      800,
			Height:     800,
			Fullscreen: false},
		LoadingScene: &scene.LoadingScene{},
		FirstScene:   &scene.GameScene{},
	})
}
