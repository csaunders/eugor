package lighting

import (
	"eugor/algebra"
	"eugor/dungeon"
	"github.com/nsf/termbox-go"
)

type MapMemory struct {
	RememberedLocations [][]bool
}

func makeMapMemory(d *dungeon.TileMap) *MapMemory {
	locations := make([][]bool, d.Width)
	for i := range locations {
		locations[i] = make([]bool, d.Height)
	}
	return &MapMemory{RememberedLocations: locations}
}

func (m *MapMemory) withinBounds(x, y int) bool {
	if x < 0 || x >= len(m.RememberedLocations) {
		return false
	}
	if y < 0 || y >= len(m.RememberedLocations[x]) {
		return false
	}
	return true
}

func (m *MapMemory) remember(x, y int) {
	if m.withinBounds(x, y) {
		m.RememberedLocations[x][y] = true
	}
}

func (m *MapMemory) isRemembered(x, y int) bool {
	if m.withinBounds(x, y) {
		return m.RememberedLocations[x][y]
	}
	return false
}

var memory *MapMemory

func ApplyFog(p algebra.Point, d *dungeon.TileMap, lights []Lightsource) {
	if memory == nil {
		memory = makeMapMemory(d)
	}
	w, h := termbox.Size()
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if withinLight(x, y, p.X, p.Y, lights) {
				markSeen(x, y, d)
				continue
			} else if memory.isRemembered(x, y) {
				tile := d.FetchTile(x, y)
				termbox.SetCell(x, y, tile.Char, termbox.ColorWhite, termbox.ColorBlack)
			} else {
				termbox.SetCell(x, y, '.', termbox.ColorWhite, termbox.ColorBlack)
			}
		}
	}
}

func withinLight(x, y, xOffset, yOffset int, lights []Lightsource) bool {
	for _, light := range lights {
		actualX, actualY := x, y
		if light.Projection() == Relative {
			actualX += xOffset
			actualY += yOffset
		}

		if light.IsLighting(actualX, actualY) {
			return true
		}
	}
	return false
}

func markSeen(x int, y int, d *dungeon.TileMap) {
	tile := d.FetchTile(x, y)
	if tile != nil && !tile.Walkable {
		memory.remember(x, y)
	}
}
