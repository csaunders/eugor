package main

import (
	"eugor/camera"
	"eugor/dungeon"
	"eugor/lighting"
	"eugor/logger"
	"eugor/sprites"
	"fmt"
	"github.com/nsf/termbox-go"
)

func main() {
	fog, running := true, true
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	torch1 := lighting.NewTorch(23, 26).Tick()
	torch2 := lighting.NewTorch(47, 20).Tick()

	char := sprites.MakeCharacter(8, 12, termbox.ColorMagenta)
	log := logger.Logger{Render: false}
	// width, height := termbox.Size()
	maze := dungeon.LoadTilemap("./empty.tlm")
	// maze := dungeon.NewTileMap(width, height)
	// maze = maze.AddRoom(5, 10, 10, 15)
	// maze = maze.AddRoom(20, 20, 20, 10)
	// maze = maze.AddRoom(14, 21, 7, 3)
	// maze = maze.AddRoom(41, 0, width-41, height)
	// maze = maze.AddDoor(39, 26)
	// maze = maze.AddDoor(41, 26)
	// maze = maze.AddDoor(14, 22)
	// maze = maze.AddDoor(20, 22)

	for running {
		termbox.Clear(termbox.ColorGreen, termbox.ColorBlack)
		// maze = maze.AdjustCamera(char.X(), char.Y())
		// char.DrawInCenter = maze.IsOffset()
		//maze.Draw()
		//char.Draw()
		characterFocus, dungeonStartPoint, meta := camera.CameraDraw(maze, char)
		lights := []lighting.Lightsource{char.Vision(), torch1, torch2}
		if fog {
			dungeon.ApplyFog(maze, lights)
		}
		log.Draw()
		termbox.Flush()
		torch1 = torch1.Tick()
		torch2 = torch2.Tick()
		event := termbox.PollEvent()
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
			event := logger.Event{LogLevel: logger.Info, Message: fmt.Sprintf("Character Position: (%d, %d)", char.X(), char.Y())}
			log = log.AppendEvent(event)
		case event.Ch == 'f':
			fog = !fog
		case event.Ch == 's':
			x, y := termbox.Size()
			event := logger.Event{LogLevel: logger.Info, Message: fmt.Sprintf("Screen Size: (%d, %d)", x, y)}
			log = log.AppendEvent(event)
		case event.Ch == 'm':
			event := logger.Event{LogLevel: logger.Info, Message: fmt.Sprintf("(%s)Character Draw Position: (%d, %d)\tDungeon Start Point: (%d, %d)", meta, characterFocus.X, characterFocus.Y, dungeonStartPoint.X, dungeonStartPoint.Y)}
			log = log.AppendEvent(event)
		case event.Ch == '`':
			log = log.ToggleRender()
		default:
			termbox.SetCell(10, 10, event.Ch, termbox.ColorRed, termbox.ColorBlack)
		}
	}
}
