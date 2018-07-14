package scene

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/dt-rush/sameriver/engine"
)

func (s *GameScene) Name() string {
	return "game-scene"
}

func (s *GameScene) Update(dt_ms float64) {
	s.w.Update(dt_ms)
}

func (s *GameScene) Draw(window *sdl.Window, renderer *sdl.Renderer) {
	// draw the score
	renderer.Copy(s.scoreTexture, nil, &s.scoreRect)
	// draw the player
	playerPos := &s.w.Em.Components.Position[s.player.ID]
	playerBox := &s.w.Em.Components.Box[s.player.ID]
	renderer.SetDrawColor(255, 255, 255, 255)
	s.game.Screen.FillRect(renderer, playerPos, playerBox)
}

func (s *GameScene) HandleKeyboardState(kb []uint8) {
	s.playerHandleKeyboardState(kb)
}

func (s *GameScene) HandleKeyboardEvent(ke *sdl.KeyboardEvent) {
	if ke.Type == sdl.KEYDOWN {
		if ke.Keysym.Sym == sdl.K_s {
			s.score++
			s.updateScoreTexture()
		}
	}
}

func (s *GameScene) IsDone() bool {
	return s.ended
}

func (s *GameScene) NextScene() engine.Scene {
	return nil
}

func (s *GameScene) IsTransient() bool {
	return true
}

func (s *GameScene) Destroy() {
	if !s.destroyed {
		s.destroyed = true
		sdl.Do(func() {
			s.scoreFont.Close()
			s.scoreSurface.Free()
			s.scoreTexture.Destroy()
		})
	}
}
