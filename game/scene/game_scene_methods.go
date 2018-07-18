package scene

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/dt-rush/sameriver/engine"

	"github.com/dt-rush/space-cats/game/systems"
)

func (s *GameScene) buildWorld() {
	// construct world object
	s.w = engine.NewWorld(s.game.WindowSpec.Width, s.game.WindowSpec.Height)
	// add systems
	s.w.AddSystems(
		engine.NewPhysicsSystem(),
		engine.NewSpatialHashSystem(32, 32),
		engine.NewCollisionSystem(100*time.Millisecond),
		systems.NewCoinDespawnAtEdgeSystem(),
	)
	// get updated entity list of coins
	s.coins = s.w.EntitiesWithTag("coin")
	// add spawn random coin logic
	s.w.AddWorldLogic("spawn-random-coin", s.spawnRandomCoin)
	// add player coin collision logic
	s.w.AddWorldLogic("player-collect-coin", s.playerCollectCoin)
	// activate all world logics
	s.w.ActivateAllWorldLogics()
}

func (s *GameScene) spawnInitialEntities() {
	var err error
	mass := 1.0
	s.player, err = s.w.SpawnUnique(
		"player",
		[]string{},
		engine.ComponentSet{
			Position: &engine.Vec2D{50, 50},
			Velocity: &engine.Vec2D{0, 0},
			Box:      &engine.Vec2D{2, 2},
			Mass:     &mass,
		},
	)
	if err != nil {
		panic(err)
	}
}

func (s *GameScene) SimpleEntityDraw(
	r *sdl.Renderer, e *engine.Entity, c sdl.Color) {

	box := e.GetBox()
	pos := e.GetPosition().ShiftedCenterToBottomLeft(box)
	r.SetDrawColor(c.R, c.G, c.B, c.A)
	s.game.Screen.FillRect(r, &pos, box)
}

func (s *GameScene) playerHandleKeyboardState(kb []uint8) {
	v := s.player.GetVelocity()
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
	s.scoreSurface, err = s.UIFont.RenderUTF8Solid(
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
	w, h, err := s.UIFont.SizeUTF8(score_msg)
	if err != nil {
		panic(err)
	}
	s.scoreRect = sdl.Rect{10, 10, int32(w), int32(h)}
}

func (s *GameScene) spawnRandomCoin() {
	if rand.Float64() < 0.8 && s.w.EntitiesWithTag("coin").Length() < 1000 {
		mass := 1.0
		c, err := s.w.Spawn(
			[]string{"coin"},
			engine.ComponentSet{
				Position: &engine.Vec2D{
					rand.Float64()*float64(s.w.Width/3) + s.w.Width/3,
					rand.Float64()*float64(s.w.Height/3) + s.w.Height/3,
				},
				Velocity: &engine.Vec2D{0, 0},
				Box:      &engine.Vec2D{4, 4},
				Mass:     &mass,
			},
		)
		if err != nil {
			fmt.Println(err)
			return
		}
		s.w.AddEntityLogic(c, s.coinLogic(c))
		s.w.ActivateEntityLogic(c)
	}
}

func (s *GameScene) coinLogic(c *engine.Entity) func() {
	return func() {
		dist := c.GetPosition().Sub(*s.player.GetPosition())
		*c.GetVelocity() = dist.Unit().Scale(0.1 * (1.0 - dist.Magnitude()/float64(s.w.Width)))
	}
}

func (s *GameScene) playerCollectCoin() {
	if s.playerCoinCollision == nil {
		s.subscribeToPlayerCoinCollision()
	}
	for len(s.playerCoinCollision.C) > 0 {
		e := <-s.playerCoinCollision.C
		coin := e.Data.(engine.CollisionData).Other
		s.w.Despawn(coin)
		s.augmentScore(10)
		s.growPlayer(0.5)
	}
}

func (s *GameScene) subscribeToPlayerCoinCollision() {
	s.playerCoinCollision = s.w.Events.Subscribe(
		"player-hit-coin",
		engine.CollisionEventFilter(
			func(c engine.CollisionData) bool {
				return c.This == s.player &&
					c.Other.GetTagList().Has("coin")
			}),
	)
}

func (s *GameScene) augmentScore(x int) {
	s.score += x
	s.updateScoreTexture()
}

func (s *GameScene) growPlayer(increase float64) {
	playerBox := s.player.GetBox()
	if playerBox.X < 50 && playerBox.Y < 50 {
		playerBox.X += increase
		playerBox.Y += increase
	}
}
