package dungeon

import (
	"eugor/algebra"
	"eugor/lighting"
	"strconv"
	"strings"
)

type MapData struct {
	Maze        TileMap
	PlayerStart algebra.Point
	MazeLights  []lighting.Lightsource
}

type LayerInformation struct {
	Type string
	Data []string
}

func TileMapToLayer(t TileMap) LayerInformation {
	data := make([]string, t.Height)
	placeholder := make([]string, t.Width)
	for y, row := range t.Tiles {
		for x, item := range row {
			placeholder[x] = strconv.FormatUint(uint64(item), 10)
		}
		data[y] = strings.Join(placeholder, ",")
	}
	return LayerInformation{Type: "map", Data: data}
}
