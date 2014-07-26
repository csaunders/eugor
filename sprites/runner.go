package sprites

import (
	"eugor"
	"eugor/dungeon"
)

type RunnerLogic struct {
	maze *dungeon.TileMap
}

func MakeRunner(d *dungeon.TileMap) RunnerLogic {
	return RunnerLogic{maze: d}
}

func (r RunnerLogic) Move(p, player eugor.Point) eugor.Point {
	x := p.X
	y := p.Y
	if p.Distance(player) > 3 {
		return p
	}
	diff := p.Minus(player)
	if diff.X < 0 {
		x -= 1
	} else if diff.X > 0 {
		x += 1
	}
	if diff.Y < 0 {
		y -= 1
	} else if diff.Y > 0 {
		y += 1
	}
	if r.maze.CanMoveTo(x, y) {
		return eugor.Point{X: x, Y: y}
	}
	return p
}
