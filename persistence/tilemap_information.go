package persistence

import (
	"eugor"
	"eugor/dungeon"
	"eugor/lighting"
	"fmt"
	"strings"
)

type MapData struct {
	Maze        *dungeon.TileMap
	PlayerStart eugor.Point
	MazeLights  []lighting.Lightsource
}

type LayerInformation struct {
	Type string
	Data []string
}

func TileMapToLayer(t *dungeon.TileMap) LayerInformation {
	data := make([]string, t.Height)
	placeholder := make([]string, t.Width)
	for y := 0; y < t.Height; y++ {
		for x := 0; x < t.Width; x++ {
			item := t.Tiles[x][y]
			placeholder[x] = fmt.Sprintf("%d", item)
		}
		data[y] = strings.Join(placeholder, ",")
	}

	return LayerInformation{Type: "map", Data: data}
}
