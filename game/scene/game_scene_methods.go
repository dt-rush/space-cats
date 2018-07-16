package scene

import (
	"fmt"
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/dt-rush/sameriver/engine"
)

func (s *GameScene) buildWorld() {
	// construct world object
	s.w = engine.NewWorld(s.game.WindowSpec.Width, s.game.WindowSpec.Height)
	// add systems
	s.w.AddSystems(
		engine.NewPhysicsSystem(),
		engine.NewSpatialHashSystem(16, 16),
		engine.NewCollisionSystem(),
	)
	// get updated entity list of coins
	s.coins = s.w.Em.EntitiesWithTag("coin")
	// add spawn random coin logic
	s.w.AddWorldLogic("spawn-random-coin", s.spawnRandomCoinLogic)
	// subscribe to player coin collision events
	s.playerCoinCollision = s.w.Ev.Subscribe(
		"player-hit-coin",
		engine.CollisionEventFilter(
			func(c engine.CollisionData) bool {
				return c.EntityA == s.player &&
					s.w.Em.EntityHasTag(c.EntityB, "coin")
			}),
	)
	// add player coin collision logic
	s.w.AddWorldLogic("player-collect-coin", s.playerCollectCoinLogic)
	// activate all world logics
	s.w.ActivateAllWorldLogics()
}

func (s *GameScene) spawnInitialEntities() {
	var err error
	mass := 1.0
	s.player, err = s.w.Em.SpawnUnique(
		"player",
		engine.SpawnRequestData{
			Components: engine.ComponentSet{
				Position: &engine.Vec2D{50, 50},
				Velocity: &engine.Vec2D{0, 0},
				Box:      &engine.Vec2D{20, 20},
				Mass:     &mass,
			},
		})
	if err != nil {
		panic(err)
	}
}

func (s *GameScene) SimpleEntityDraw(
	r *sdl.Renderer, e *engine.EntityToken, c sdl.Color) {

	pos := &s.w.Em.Components.Position[e.ID]
	box := &s.w.Em.Components.Box[e.ID]
	r.SetDrawColor(c.R, c.G, c.B, c.A)
	s.game.Screen.FillRect(r, pos, box)
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

func (s *GameScene) spawnRandomCoinLogic() {
	if rand.Float64() < 0.8 && s.w.Em.EntitiesWithTag("coin").Length() < 1000 {
		s.spawnRandomCoin()
	}
}

func (s *GameScene) spawnRandomCoin() {
	mass := 1.0
	_, err := s.w.Em.Spawn(engine.SpawnRequestData{
		Components: engine.ComponentSet{
			Position: &engine.Vec2D{
				rand.Float64() * float64(s.w.Width),
				rand.Float64() * float64(s.w.Height),
			},
			Velocity: &engine.Vec2D{0, 0},
			Box:      &engine.Vec2D{4, 4},
			Mass:     &mass,
		},
		Tags: []string{"coin"},
	})
	if err != nil {
		fmt.Println(err)
	}
}

func (s *GameScene) playerCollectCoinLogic() {
	for len(s.playerCoinCollision.C) > 0 {
		e := <-s.playerCoinCollision.C
		c := e.Data.(engine.CollisionData)
		s.score += 10
		s.updateScoreTexture()
		coin := c.EntityB
		s.w.Em.Despawn(coin)
		playerBox := &s.w.Em.Components.Box[s.player.ID]
		if playerBox.X < 50 && playerBox.Y < 50 {
			playerBox.X += 2
			playerBox.Y += 2
		}
	}
}
