package main

import (
	"eugor/algebra"
	"eugor/camera"
	"eugor/dungeon"
	"eugor/lighting"
	"eugor/logger"
	"eugor/persistence"
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

	mapConfiguration := persistence.LoadTilemap("./persisted.tlm")
	maze := mapConfiguration.Maze
	lights := mapConfiguration.MazeLights

	start := mapConfiguration.PlayerStart

	char := sprites.MakeCharacter(start.X, start.Y, termbox.ColorMagenta)

	dungeonMaster := sprites.MakeDungeonMaster(char, maze)

	log := logger.Logger{Render: false}

	mapContext := sprites.DefaultMapContext(maze)

	updateWorld := func() {
		for _, light := range lights {
			light.Tick()
		}
		dungeonMaster.Tick(char.Position())
	}

	for running {
		termbox.Clear(termbox.ColorGreen, termbox.ColorBlack)
		characterFocus, dungeonStartPoint, meta := camera.CameraDraw(maze, char, dungeonMaster.Drawables())
		if fog {
			newLights := []lighting.Lightsource{char.Vision(characterFocus, maze)}
			// lighting.ApplyFog(dungeonStartPoint, maze, append(lights, char.Vision(characterFocus, maze)))
			lighting.ApplyFog(dungeonStartPoint, maze, newLights)
		}
		log.Draw()
		mapContext.Draw()
		termbox.Flush()
		event := termbox.PollEvent()
		charPoint := algebra.MakePoint(char.X(), char.Y())
		switch {
		case event.Key == termbox.KeyEsc:
			running = false
		case event.Ch == 'i':
			mapContext.Toggle(charPoint)
		case mapContext.IsFocused():
			mapContext.HandleInput(charPoint, event)
		case char.IsMovementEvent(event):
			x, y := char.PredictedMovement(event.Key)
			if maze.CanMoveTo(x, y) && !dungeonMaster.Occupied(x, y) {
				char.Move(event.Key)
			} else if maze.CanMoveTo(x, y) && dungeonMaster.Occupied(x, y) {
				didHit := dungeonMaster.Interact(x, y, char)
				if didHit {
					log = log.AppendEvent(logger.Event{LogLevel: logger.Info, Message: "Creature has been defeated!"})
				} else {
					log = log.AppendEvent(logger.Event{LogLevel: logger.Info, Message: "Womp Womp, you missed :'("})
				}
			} else if maze.CanInteractWith(x, y) {
				maze.Interact(x, y)
			}
			updateWorld()
		case event.Ch == 'l':
			event := logger.Event{LogLevel: logger.Info, Message: fmt.Sprintf("Character Position: (%d, %d)", char.X(), char.Y())}
			log = log.AppendEvent(event)
		case event.Ch == 'f':
			fog = !fog
		case event.Ch == 's':
			x, y := termbox.Size()
			event := logger.Event{LogLevel: logger.Info, Message: fmt.Sprintf("Screen Size: (%d, %d)", x, y)}
			log = log.AppendEvent(event)
		case event.Ch == 'S':
			persistMapDetails(maze, char, lights)
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

func persistMapDetails(maze *dungeon.TileMap, player *sprites.Character, lights []lighting.Lightsource) {
	start := algebra.MakePoint(player.X(), player.Y())
	data := persistence.MapData{Maze: maze, PlayerStart: start, MazeLights: lights}
	persistence.SaveTilemap(data, "persisted.tlm")
}
