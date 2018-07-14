/*
 *
 *
 *
 *
**/

package scene

import (
	"github.com/dt-rush/sameriver/engine"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type GameScene struct {

	// Scene "abstract class members"

	// whether the scene is running
	ended bool
	// used to make destroy() idempotent
	destroyed bool
	// the game
	game *engine.Game

	// GameScene members
	w *engine.World

	// score of player in this scene
	score     int
	scoreFont *ttf.Font
	// surface used to build texture
	scoreSurface *sdl.Surface
	// texture of the above, for Renderer.Copy() in draw()
	scoreTexture *sdl.Texture
	// score texture screen width
	scoreRect sdl.Rect
}
