package scene

import (
	"fmt"
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/dt-rush/sameriver/engine"
)

func (s *GameScene) buildWorld() {
	s.w = engine.NewWorld(s.game.WindowSpec.Width, s.game.WindowSpec.Height)
	s.w.AddSystems(
		engine.NewPhysicsSystem(),
		engine.NewSpatialHashSystem(10, 10),
		engine.NewCollisionSystem(),
	)
	s.w.AddWorldLogic("spawn-random-coin", s.spawnRandomCoinLogic)
	s.playerCoinCollision = s.w.Ev.Subscribe(
		"player-hit-coin",
		engine.NewPredicateEventQuery(
			engine.COLLISION_EVENT,
			func(e engine.Event) bool {
				c := e.Data.(engine.CollisionData)
				return ((c.EntityA == s.player &&
					s.w.Em.EntityHasTag(c.EntityB, "coin")) ||
					(c.EntityB == s.player &&
						s.w.Em.EntityHasTag(c.EntityA, "coin")))
			}),
	)
	s.w.AddWorldLogic("player-collect-coin", s.playerCollectCoinLogic)
	s.w.ActivateAllWorldLogics()
	s.coins = s.w.Em.EntitiesWithTag("coin")
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
				Box:      &engine.Vec2D{10, 10},
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

func (s *GameScene) spawnRandomCoinLogic() {
	if rand.Float64() < 0.008 {
		s.spawnRandomCoin()
	}
}

func (s *GameScene) spawnRandomCoin() {
	var err error
	mass := 1.0
	_, err = s.w.Em.Spawn(engine.SpawnRequestData{
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
		if s.w.Em.EntityHasTag(c.EntityA, "coin") {
			s.w.Em.Despawn(c.EntityA)
		}
		if s.w.Em.EntityHasTag(c.EntityB, "coin") {
			s.w.Em.Despawn(c.EntityB)
		}
	}
}
