package scene

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"

	"github.com/dt-rush/sameriver/v2"
)

type GameScene struct {

	// Scene "abstract class members"

	// whether the scene is running
	ended bool
	// used to make destroy() idempotent
	destroyed bool
	// the game
	game *sameriver.Game

	// GameScene members
	w                   *sameriver.World
	player              *sameriver.Entity
	coins               *sameriver.UpdatedEntityList
	playerCoinCollision *sameriver.EventChannel

	// for writing text to the screen
	UIFont *ttf.Font

	// score of player in this scene
	score int
	// surface used to build texture
	scoreSurface *sdl.Surface
	// texture of the above, for Renderer.Copy() in draw()
	scoreTexture *sdl.Texture
	// score texture screen width
	scoreRect sdl.Rect
}
