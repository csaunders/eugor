package sprites

import (
	"eugor/algebra"
	"github.com/nsf/termbox-go"
)

type CreatureLogic interface {
	Move(position algebra.Point, player algebra.Point) algebra.Point
}

type Creature struct {
	Position algebra.Point
	Color    termbox.Attribute
	Icon     rune
	Ai       CreatureLogic
}

func MakeCreature(x, y int, c termbox.Attribute, r rune) *Creature {
	p := algebra.MakePoint(x, y)
	return &Creature{Position: p, Color: c, Icon: r, Ai: DumbAi{}}
}

func (c *Creature) Tick(playerPosition algebra.Point) {
	c.Position = c.Ai.Move(c.Position, playerPosition)
}

func (c *Creature) X() int {
	return c.Position.X
}

func (c *Creature) Y() int {
	return c.Position.Y
}

func (c *Creature) DrawProjection(screenX, screenY, positionX, positionY int) {
	termbox.SetCell(screenX, screenY, c.Icon, c.Color, termbox.ColorBlack)
}

type DumbAi struct{}

func (d DumbAi) Move(p, player algebra.Point) algebra.Point {
	if p.X%2 == 0 {
		p.X += 1
	} else {
		p.X -= 1
	}
	return p
}
