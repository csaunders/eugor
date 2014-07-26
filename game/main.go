package main

import (
	"eugor"
	"eugor/dungeon"
	"eugor/lighting"
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

	mapContext := sprites.DefaultMapContext(maze)

	updateWorld := func() {
		for _, light := range lights {
			light.Tick()
		}
		dungeonMaster.Tick(char.Position())
	}

	for running {
		termbox.Clear(termbox.ColorGreen, termbox.ColorBlack)
		characterFocus, dungeonStartPoint, meta := eugor.CameraDraw(maze, char, dungeonMaster.Drawables())
		if fog {
			newLights := []lighting.Lightsource{char.Vision(characterFocus, maze)}
			// lighting.ApplyFog(dungeonStartPoint, maze, append(lights, char.Vision(characterFocus, maze)))
			lighting.ApplyFog(dungeonStartPoint, maze, newLights)
		}
		eugor.GlobalLog.Draw()
		mapContext.Draw()
		termbox.Flush()
		event := termbox.PollEvent()
		charPoint := eugor.MakePoint(char.X(), char.Y())
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
					eugor.GlobalLog.AppendEvent(eugor.Event{LogLevel: eugor.Info, Message: "Creature has been defeated!"})
				} else {
					eugor.GlobalLog.AppendEvent(eugor.Event{LogLevel: eugor.Info, Message: "Womp Womp, you missed :'("})
				}
			} else if maze.CanInteractWith(x, y) {
				maze.Interact(x, y)
			}
			updateWorld()
		case event.Ch == 'l':
			event := eugor.Event{LogLevel: eugor.Info, Message: fmt.Sprintf("Character Position: (%d, %d)", char.X(), char.Y())}
			eugor.GlobalLog.AppendEvent(event)
		case event.Ch == 'f':
			fog = !fog
		case event.Ch == 's':
			x, y := termbox.Size()
			event := eugor.Event{LogLevel: eugor.Info, Message: fmt.Sprintf("Screen Size: (%d, %d)", x, y)}
			eugor.GlobalLog.AppendEvent(event)
		case event.Ch == 'S':
			persistMapDetails(maze, char, lights)
		case event.Ch == 'm':
			event := eugor.Event{LogLevel: eugor.Info, Message: fmt.Sprintf("(%s)Character Draw Position: (%d, %d)\tDungeon Start Point: (%d, %d)", meta, characterFocus.X, characterFocus.Y, dungeonStartPoint.X, dungeonStartPoint.Y)}
			eugor.GlobalLog.AppendEvent(event)
		case event.Ch == '`':
			eugor.GlobalLog.ToggleRender()
		default:
			termbox.SetCell(10, 10, event.Ch, termbox.ColorRed, termbox.ColorBlack)
		}
	}
}

func persistMapDetails(maze *dungeon.TileMap, player *sprites.Character, lights []lighting.Lightsource) {
	start := eugor.MakePoint(player.X(), player.Y())
	data := persistence.MapData{Maze: maze, PlayerStart: start, MazeLights: lights}
	persistence.SaveTilemap(data, "persisted.tlm")
}
