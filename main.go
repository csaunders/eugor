package main

import (
	"eugor/logger"
	"github.com/nsf/termbox-go"
)

type Character struct {
	x int
	y int
}

func (c Character) draw() {
	termbox.SetCell(c.x, c.y, '@', termbox.ColorMagenta, termbox.ColorBlack)
}

func (c Character) move(k termbox.Key) Character {
	switch {
	case k == termbox.KeyArrowUp:
		c.y -= 1
	case k == termbox.KeyArrowDown:
		c.y += 1
	case k == termbox.KeyArrowLeft:
		c.x -= 1
	case k == termbox.KeyArrowRight:
		c.x += 1
	}
	return c
}

func (c Character) isMovementEvent(e termbox.Event) bool {
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

func main() {
	running := true
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	char := Character{x: 5, y: 5}
	logger := logger.Logger{}

	for running {
		termbox.Clear(termbox.ColorGreen, termbox.ColorBlack)
		logger.Draw()
		char.draw()
		termbox.Flush()
		event := termbox.PollEvent()
		logger.Append(event)
		switch {
		case event.Key == termbox.KeyEsc:
			running = false
		case char.isMovementEvent(event):
			char = char.move(event.Key)
		default:
			termbox.SetCell(10, 10, event.Ch, termbox.ColorRed, termbox.ColorBlack)
		}
	}
}
