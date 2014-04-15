package dungeon

import (
	"eugor/algebra"
	"eugor/lighting"
	"github.com/nsf/termbox-go"
)

func ApplyFog(p algebra.Point, d TileMap, lights []lighting.Lightsource) {
	w, h := termbox.Size()
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if withinLight(x, y, p.X, p.Y, lights) {
				continue
			} else {
				termbox.SetCell(x, y, '.', termbox.ColorWhite, termbox.ColorBlack)
			}
		}
	}
}

func withinLight(x, y, xOffset, yOffset int, lights []lighting.Lightsource) bool {
	for _, light := range lights {
		actualX, actualY := x, y
		if light.Projection() == lighting.Relative {
			actualX += xOffset
			actualY += yOffset
		}

		if light.IsLighting(actualX, actualY) {
			return true
		}
	}
	return false
}
