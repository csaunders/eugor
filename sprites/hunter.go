package sprites

import (
	"eugor"
	"eugor/dungeon"
	"eugor/lighting"
)

type HunterLogic struct {
	caster *lighting.Raycaster
}

func MakeHunter(maze *dungeon.TileMap) HunterLogic {
	caster := lighting.MakeRaycaster(maze)
	return HunterLogic{caster: caster}
}

func (h HunterLogic) Move(p, player eugor.Point) eugor.Point {
	h.caster.FlushOverlay()
	h.caster.DoLine(p.X, p.Y, player.X, player.Y)
	if h.caster.IsLighting(player.X, player.Y) {
		return h.moveToward(p, player)
	}
	return p
}

func (h HunterLogic) moveToward(p, player eugor.Point) eugor.Point {
	x, y := p.X, p.Y
	if p.X < player.X {
		x += 1
	} else if p.X > player.X {
		x -= 1
	}

	if p.Y < player.Y {
		y += 1
	} else if p.Y > player.Y {
		y -= 1
	}

	if x == player.X && y == player.Y {
		return p
	}
	return eugor.MakePoint(x, y)
}
