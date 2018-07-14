package scene

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

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
