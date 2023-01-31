package scene

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/dt-rush/sameriver/v2"
)

func (s *GameScene) Name() string {
	return "game-scene"
}

func (s *GameScene) Update(dt_ms float64, allowance_ms float64) {
	s.w.Update(allowance_ms)
}

func (s *GameScene) Draw(w *sdl.Window, r *sdl.Renderer) {
	// draw the score
	r.Copy(s.scoreTexture, nil, &s.scoreRect)
	// draw the coins
	for _, coin := range s.coins.GetEntities() {
		s.SimpleEntityDraw(r, coin, sdl.Color{255, 255, 0, 255})
	}
	// draw the player
	s.SimpleEntityDraw(r, s.player, sdl.Color{255, 255, 255, 255})
}

func (s *GameScene) HandleKeyboardState(kb []uint8) {
	s.playerHandleKeyboardState(kb)
}

func (s *GameScene) HandleKeyboardEvent(ke *sdl.KeyboardEvent) {
	if ke.Type == sdl.KEYDOWN {
		if ke.Keysym.Sym == sdl.K_SPACE {
			fmt.Println("you pressed space")
		}
	}
}

func (s *GameScene) IsDone() bool {
	return s.ended
}

func (s *GameScene) NextScene() sameriver.Scene {
	return nil
}

func (s *GameScene) End() {
	fmt.Println(s.w.DumpStatsString())
}

func (s *GameScene) IsTransient() bool {
	return true
}

func (s *GameScene) Destroy() {
	if !s.destroyed {
		s.destroyed = true
		sdl.Do(func() {
			s.UIFont.Close()
			s.scoreSurface.Free()
			s.scoreTexture.Destroy()
		})
	}
}
