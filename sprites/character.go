package sprites

import (
	"eugor/lighting"
	"github.com/nsf/termbox-go"
)

type Character struct {
	X            int
	Y            int
	DrawInCenter bool
	Color        termbox.Attribute
}

func (c Character) Draw() {
	x, y := c.X, c.Y
	if c.DrawInCenter {
		sx, sy := termbox.Size()
		x, y = sx/2, sy/2
	}
	termbox.SetCell(x, y, '@', c.Color, termbox.ColorBlack)
}

func (c Character) PredictedMovement(k termbox.Key) (int, int) {
	x, y := c.X, c.Y
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

func (c Character) Move(k termbox.Key) Character {
	c.X, c.Y = c.PredictedMovement(k)
	return c
}

func (c Character) Vision() lighting.Lightsource {
	return lighting.NewVision(c.X, c.Y, 3)
}

func (c Character) IsMovementEvent(e termbox.Event) bool {
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
