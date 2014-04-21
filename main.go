package main

import (
	"eugor/algebra"
	"eugor/camera"
	"eugor/dungeon"
	"eugor/lighting"
	"eugor/logger"
	"eugor/particles"
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

	closeDoor := sprites.Interactable{
		Name: "Close Door",
		Test: func(p algebra.Point, d dungeon.TileMap) bool {
			tile := d.FetchTile(p.X, p.Y)
			return d.CanInteractWith(p.X, p.Y) && tile.Name == "opendoor"
		},
		Action: func(p algebra.Point, d dungeon.TileMap) dungeon.TileMap {
			return d.Interact(p.X, p.Y)
		},
	}
	openDoor := sprites.Interactable{
		Name: "Open Door",
		Test: func(p algebra.Point, d dungeon.TileMap) bool {
			tile := d.FetchTile(p.X, p.Y)
			return d.CanInteractWith(p.X, p.Y) && tile.Name == "door"
		},
		Action: func(p algebra.Point, d dungeon.TileMap) dungeon.TileMap {
			return d.Interact(p.X, p.Y)
		},
	}

	mapConfiguration := dungeon.LoadTilemap("./persisted.tlm")
	maze := mapConfiguration.Maze
	lights := mapConfiguration.MazeLights

	// torch1 := lighting.NewTorch(23, 26).Tick()
	// torch2 := lighting.NewTorch(47, 20).Tick()
	// lights := []lighting.Lightsource{torch1, torch2}

	emmiter := particles.MakeEmmiter(algebra.MakePoint(30, 10), 5)
	start := mapConfiguration.PlayerStart

	char := sprites.MakeCharacter(start.X, start.Y, termbox.ColorMagenta)
	monsters := []sprites.Creature{
		sprites.MakeCreature(30, 15, termbox.ColorBlue, 'k'),
		sprites.MakeCreature(45, 15, termbox.ColorYellow, '%'),
		sprites.MakeCreature(30, 25, termbox.ColorGreen, '?'),
	}
	monsterDrawers := make([]camera.Drawable, len(monsters))
	for i, m := range monsters {
		monsterDrawers[i] = m
	}
	log := logger.Logger{Render: false}
	// width, height := termbox.Size()
	// maze := dungeon.NewTileMap(width, height)
	// maze = maze.AddRoom(5, 10, 10, 15)
	// maze = maze.AddRoom(20, 20, 20, 10)
	// maze = maze.AddRoom(14, 21, 7, 3)
	// maze = maze.AddRoom(41, 0, width-41, height)
	// maze = maze.AddDoor(39, 26)
	// maze = maze.AddDoor(41, 26)
	// maze = maze.AddDoor(14, 22)
	// maze = maze.AddDoor(20, 22)

	mapContext := sprites.MapContext{TileMap: maze}
	mapContext = mapContext.AddInteraction(closeDoor)
	mapContext = mapContext.AddInteraction(openDoor)

	for running {
		termbox.Clear(termbox.ColorGreen, termbox.ColorBlack)
		characterFocus, dungeonStartPoint, meta := camera.CameraDraw(maze, char, monsterDrawers)
		emmiter = emmiter.Update()
		emmiter.Draw()
		if fog {
			dungeon.ApplyFog(dungeonStartPoint, maze, append(lights, char.Vision(characterFocus)))
		}
		log.Draw()
		mapContext.Draw()
		termbox.Flush()
		for i, light := range lights {
			lights[i] = light.Tick()
		}
		for i, m := range monsters {
			monsters[i] = m.Tick()
			monsterDrawers[i] = monsters[i]
		}
		event := termbox.PollEvent()
		charPoint := algebra.MakePoint(char.X(), char.Y())
		switch {
		case event.Key == termbox.KeyEsc:
			running = false
		case event.Ch == 'i':
			mapContext = mapContext.Toggle(charPoint)
		case mapContext.IsFocused():
			mapContext = mapContext.HandleInput(charPoint, event)
		case char.IsMovementEvent(event):
			x, y := char.PredictedMovement(event.Key)
			if maze.CanMoveTo(x, y) {
				char = char.Move(event.Key)
			} else if maze.CanInteractWith(x, y) {
				maze = maze.Interact(x, y)
			}
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

func persistMapDetails(maze dungeon.TileMap, player sprites.Character, lights []lighting.Lightsource) {
	start := algebra.MakePoint(player.X(), player.Y())
	data := dungeon.MapData{Maze: maze, PlayerStart: start, MazeLights: lights}
	dungeon.SaveTilemap(data, "persisted.tlm")
}
