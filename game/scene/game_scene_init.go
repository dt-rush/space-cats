package scene

import (
	"time"

	"github.com/dt-rush/sameriver/engine"
	"github.com/veandco/go-sdl2/ttf"
)

func (s *GameScene) Init(game *engine.Game, config map[string]string) {
	var err error
	s.destroyed = false
	s.game = game

	if s.scoreFont, err = ttf.OpenFont("assets/fixedsys.ttf", 10); err != nil {
		panic(err)
	}
	s.score = 0
	s.updateScoreTexture()

	// just to play a little loading screen fun
	time.Sleep(1 * time.Second)
}