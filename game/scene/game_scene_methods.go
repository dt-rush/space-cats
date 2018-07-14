package scene

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/dt-rush/sameriver/engine"
)

func (s *GameScene) buildWorld() {
	s.w = engine.NewWorld(1024, 1024)
	s.w.AddSystems(
		engine.NewPhysicsSystem(),
	)
}

func (s *GameScene) spawnEntities() {
	var err error
	mass := 1.0
	s.player, err = s.w.Em.Spawn(engine.SpawnRequestData{
		Components: engine.ComponentSet{
			Position: &engine.Vec2D{50, 50},
			Velocity: &engine.Vec2D{0, 0},
			Box:      &engine.Vec2D{10, 10},
			Mass:     &mass,
		},
	})
	if err != nil {
		panic(err)
	}
}

func (s *GameScene) playerHandleKeyboardState(kb []uint8) {
	v := &s.w.Em.Components.Velocity[s.player.ID]
	// get player v1
	v.X = 0.2 * float64(
		int8(kb[sdl.SCANCODE_D]|kb[sdl.SCANCODE_RIGHT])-
			int8(kb[sdl.SCANCODE_A]|kb[sdl.SCANCODE_LEFT]))
	v.Y = 0.2 * float64(
		int8(kb[sdl.SCANCODE_W]|kb[sdl.SCANCODE_UP])-
			int8(kb[sdl.SCANCODE_S]|kb[sdl.SCANCODE_DOWN]))
}

func (s *GameScene) updateScoreTexture() {

	if s.scoreSurface != nil {
		s.scoreSurface.Free()
	}
	if s.scoreTexture != nil {
		s.scoreTexture.Destroy()
	}
	// render message ("press space") surface
	score_msg := fmt.Sprintf("%d", s.score)
	var err error
	s.scoreSurface, err = s.scoreFont.RenderUTF8Solid(
		score_msg,
		sdl.Color{255, 255, 255, 255})
	if err != nil {
		panic(err)
	}
	// create the texture
	s.scoreTexture, err = s.game.Renderer.CreateTextureFromSurface(s.scoreSurface)
	if err != nil {
		panic(err)
	}
	// set the width of the texture on screen
	s.scoreRect = sdl.Rect{
		10,
		10,
		int32(len(score_msg) * (s.game.WindowSpec.Width / 21)),
		20}
}
