package scene

import (
	"time"

	"github.com/dt-rush/sameriver/engine"
	"github.com/veandco/go-sdl2/ttf"
)

func (s *GameScene) Init(game *engine.Game, config map[string]string) {
	var err error
	// set scene "abstract base class" members
	s.destroyed = false
	s.game = game

	// set up score font
	if s.scoreFont, err = ttf.OpenFont("assets/fixedsys.ttf", 10); err != nil {
		panic(err)
	}
	s.score = 0
	s.updateScoreTexture()

	// build world and spawn entities
	s.buildWorld()
	s.spawnInitialEntities()

	// just to play a little loading screen fun
	time.Sleep(1 * time.Second)
}
