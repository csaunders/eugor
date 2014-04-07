package main

import (
	"eugor/dungeon"
	"eugor/logger"
	"eugor/sprites"
	"github.com/nsf/termbox-go"
)

func main() {
	running := true
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	char := sprites.Character{X: 5, Y: 5, Color: termbox.ColorMagenta}
	logger := logger.Logger{Render: false}
	dungeon := dungeon.NewMap()

	for running {
		if dungeon.IsWithinBounds(char) {
			char.Color = termbox.ColorYellow
		} else {
			char.Color = termbox.ColorRed
		}
		termbox.Clear(termbox.ColorGreen, termbox.ColorBlack)
		dungeon.Draw()
		logger.Draw()
		char.Draw()
		termbox.Flush()
		event := termbox.PollEvent()
		logger = logger.Append(event)
		switch {
		case event.Key == termbox.KeyEsc:
			running = false
		case char.IsMovementEvent(event):
			char = char.Move(event.Key)
			// dungeon = dungeon.Move(event.Key)
		case event.Ch == '`':
			logger = logger.ToggleRender()
		default:
			termbox.SetCell(10, 10, event.Ch, termbox.ColorRed, termbox.ColorBlack)
		}
	}
}
