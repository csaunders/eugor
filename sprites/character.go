package sprites

import (
	"eugor/algebra"
	"eugor/lighting"
	"github.com/nsf/termbox-go"
)

type Character struct {
	x            int
	y            int
	DrawInCenter bool
	Color        termbox.Attribute
}

func MakeCharacter(x, y int, color termbox.Attribute) *Character {
	return &Character{x: x, y: y, Color: color}
}

func (c *Character) DrawProjection(screenX, screenY, positionX, positionY int) {
	termbox.SetCell(screenX, screenY, '@', c.Color, termbox.ColorBlack)
}

func (c *Character) X() int {
	return c.x
}

func (c *Character) Y() int {
	return c.y
}

func (c *Character) Draw() {
	x, y := c.x, c.y
	if c.DrawInCenter {
		sx, sy := termbox.Size()
		x, y = sx/2, sy/2
	}
	termbox.SetCell(x, y, '@', c.Color, termbox.ColorBlack)
}

func (c *Character) PredictedMovement(k termbox.Key) (int, int) {
	x, y := c.x, c.y
	switch {
	case k == termbox.KeyArrowUp:
		y -= 1
	case k == termbox.KeyArrowDown:
		y += 1
	case k == termbox.KeyArrowLeft:
		x -= 1
	case k == termbox.KeyArrowRight:
		x += 1
	}
	return x, y
}

func (c *Character) Move(k termbox.Key) {
	c.x, c.y = c.PredictedMovement(k)
}

func (c *Character) Vision(p algebra.Point) lighting.Lightsource {
	return lighting.NewVision(p.X, p.Y, 3)
}

func (c *Character) IsMovementEvent(e termbox.Event) bool {
	validEvents := []termbox.Key{
		termbox.KeyArrowUp,
		termbox.KeyArrowDown,
		termbox.KeyArrowLeft,
		termbox.KeyArrowRight,
	}
	for _, key := range validEvents {
		if e.Key == key {
			return true
		}
	}
	return false
}
