package dungeon

import (
	"eugor/sprites"
	"github.com/nsf/termbox-go"
)

func ApplyFog(d TileMap, c sprites.Character) {
	visibility := c.Visibility()
	for x := range d.Tiles {
		for y := range d.Tiles[x] {
			dX := x - c.X
			dY := y - c.Y
			if dX >= (-visibility*2) && dX <= (visibility*2) && dY >= -visibility && dY <= visibility {
				continue
			} else {
				termbox.SetCell(x, y, '.', termbox.ColorWhite, termbox.ColorBlack)
			}
		}
	}
}
