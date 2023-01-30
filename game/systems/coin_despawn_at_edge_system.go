package systems

import (
	"github.com/dt-rush/sameriver/engine"
)

type CoinDespawnAtEdgeSystem struct {
	w     *engine.World
	sh    *engine.SpatialHashSystem `sameriver-system-dependency:"-"`
	coins *engine.UpdatedEntityList
}

func NewCoinDespawnAtEdgeSystem() *CoinDespawnAtEdgeSystem {
	return &CoinDespawnAtEdgeSystem{}
}

func (s *CoinDespawnAtEdgeSystem) LinkWorld(w *engine.World) {
	s.w = w
	s.coins = s.w.EntitiesWithTag("coin")
}

func (s *CoinDespawnAtEdgeSystem) Update() {
	for y := 0; y <= s.sh.GridY-1; y += (s.sh.GridY - 1) {
		for x := 0; x < s.sh.GridX; x++ {
			cell := s.sh.Table[x][y]
			for _, e := range cell {
				if e.GetTagList("GenericTags").Has("coin") {
					pos := e.GetVec2D("Position")
					box := e.GetVec2D("Box")
					if pos.Y < box.Y || (s.w.Height-pos.Y) < box.Y {
						s.w.Despawn(e)
					}
				}
			}
		}
	}
	for x := 0; x <= s.sh.GridX-1; x += (s.sh.GridX - 1) {
		for y := 0; y < s.sh.GridY; y++ {
			cell := s.sh.Table[x][y]
			for _, e := range cell {
				if e.GetTagList("GenericTags").Has("coin") {
					pos := e.GetVec2D("Position")
					box := e.GetVec2D("Box")
					if pos.X < box.X || (s.w.Width-pos.X) < box.X {
						s.w.Despawn(e)
					}
				}
			}
		}
	}
}

func (s *CoinDespawnAtEdgeSystem) GetComponentDeps() []string {
	return []string{"TagList,GenericTags", "Vec2D,Position", "Vec2D,Box"}
}
