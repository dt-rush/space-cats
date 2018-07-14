package scene

import (
	"github.com/dt-rush/sameriver/engine"
	"github.com/dt-rush/sameriver/engine/utils"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func (s *LoadingScene) Init(game *engine.Game, config map[string]string) {
	var err error
	s.game = game
	if !s.initialized {
		s.destroyed = false
		s.accum_5000 = utils.CreateTimeAccumulator(5000)
		s.message_font, err = ttf.OpenFont("./assets/fixedsys.ttf", 10)
		if err != nil {
			panic(err)
		}
		s.message_surface, err = s.message_font.RenderUTF8Solid("Loading",
			sdl.Color{255, 255, 255, 255})
		if err != nil {
			panic(err)
		}
		s.message_texture, err = s.game.Renderer.CreateTextureFromSurface(
			s.message_surface)
		if err != nil {
			panic(err)
		}
		s.initialized = true
	}
}
