package sprites

import (
	"eugor"
	"eugor/dungeon"
	"math/rand"
	"time"
)

type WalkerLogic struct {
	maze *dungeon.TileMap
	prng *rand.Rand
}

func MakeWalker(maze *dungeon.TileMap) WalkerLogic {
	prng := rand.New(rand.NewSource(time.Now().UnixNano()))
	return WalkerLogic{maze: maze, prng: prng}
}

func (w WalkerLogic) Move(p, player eugor.Point) eugor.Point {
	x := w.generateAdjustment()
	y := w.generateAdjustment()
	newP := eugor.MakePoint(p.X+x, p.Y+y)
	if w.maze.CanMoveTo(newP.X, newP.Y) {
		return newP
	}
	return p
}

func (w WalkerLogic) generateAdjustment() int {
	switch w.prng.Intn(3) {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return -1
	}
}
