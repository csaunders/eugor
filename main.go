package main

import (
	"eugor/dungeon"
	"eugor/logger"
	"eugor/sprites"
	"fmt"
	"github.com/nsf/termbox-go"
)

func main() {
	running := true
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	char := sprites.Character{X: 8, Y: 12, Color: termbox.ColorMagenta}
	log := logger.Logger{Render: false}
	width, height := termbox.Size()
	maze := dungeon.NewTileMap(width, height)
	maze = maze.AddRoom(5, 10, 10, 15)
	maze = maze.AddRoom(20, 20, 20, 10)
	maze = maze.AddRoom(14, 21, 7, 3)
	maze = maze.AddRoom(41, 0, width-41, height)
	maze = maze.AddDoor(39, 26)
	maze = maze.AddDoor(41, 26)
	maze = maze.AddDoor(14, 22)
	maze = maze.AddDoor(20, 22)

	for running {
		termbox.Clear(termbox.ColorGreen, termbox.ColorBlack)
		maze.Draw()
		log.Draw()
		char.Draw()
		dungeon.ApplyFog(maze, char)
		termbox.Flush()
		event := termbox.PollEvent()
		log = log.Append(event)
		switch {
		case event.Key == termbox.KeyEsc:
			running = false
		case char.IsMovementEvent(event):
			x, y := char.PredictedMovement(event.Key)
			if maze.CanMoveTo(x, y) {
				char = char.Move(event.Key)
			} else if maze.CanInteractWith(x, y) {
				maze = maze.Interact(x, y)
			}
			// maze = maze.Move(event.Key)
		case event.Ch == 'l':
			event := logger.Event{LogLevel: logger.Info, Message: fmt.Sprintf("Character Position: (%d, %d)", char.X, char.Y)}
			log = log.AppendEvent(event)
		case event.Ch == '`':
			log = log.ToggleRender()
		default:
			termbox.SetCell(10, 10, event.Ch, termbox.ColorRed, termbox.ColorBlack)
		}
	}
}
