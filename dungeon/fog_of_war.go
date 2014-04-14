package dungeon

import (
	"eugor/lighting"
	"github.com/nsf/termbox-go"
)

func ApplyFog(d TileMap, lights []lighting.Lightsource) {
	for x := range d.Tiles {
		for y := range d.Tiles[x] {
			if withinLight(x, y, lights) {
				continue
			} else {
				termbox.SetCell(x, y, '.', termbox.ColorWhite, termbox.ColorBlack)
			}
		}
	}
}

func withinLight(x, y int, lights []lighting.Lightsource) bool {
	for _, light := range lights {
		if light.IsLighting(x, y) {
			return true
		}
		// dX := x - light.X()
		// dY := y - light.Y()
		// visibility := light.Intensity()
		// if dX >= (-visibility*2) && dX <= (visibility*2) && dY >= -visibility && dY <= visibility {
		// 	return true
		// }
	}
	return false
}
