package scene

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/dt-rush/sameriver/v2"
)

func (s *LoadingScene) Name() string {
	return "loading-scene"
}

func (s *LoadingScene) Update(dt_ms float64, allowance_ms float64) {
	s.accum_5000.Tick(dt_ms)
}

func (s *LoadingScene) Draw(window *sdl.Window, renderer *sdl.Renderer) {
	W := int32(s.game.WindowSpec.Width)
	H := int32(s.game.WindowSpec.Height)
	x := W / 3
	y := float64(H*2/5) +
		float64(H/10)*math.Sin(5*2*math.Pi*s.accum_5000.Completion())
	msg_dst := sdl.Rect{
		x,
		int32(y),
		W / 3,
		20}
	renderer.Copy(s.message_texture, nil, &msg_dst)
}

func (s *LoadingScene) HandleKeyboardState(kb []uint8) {
	// null implementation
}
func (s *LoadingScene) HandleKeyboardEvent(ke *sdl.KeyboardEvent) {
	// null implementation
}

func (s *LoadingScene) IsDone() bool {
	return false
}

func (s *LoadingScene) NextScene() sameriver.Scene {
	return nil
}

func (s *LoadingScene) End() {
}

func (s *LoadingScene) IsTransient() bool {
	return false
}

func (s *LoadingScene) Destroy() {
	if !s.destroyed {
		s.destroyed = true
		sdl.Do(func() {
			s.message_font.Close()
			s.message_surface.Free()
			s.message_texture.Destroy()
		})
	}
}
