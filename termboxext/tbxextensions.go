package termboxext

import (
	"github.com/nsf/termbox-go"
)

func DrawString(x, y int, msg string, fg, bg termbox.Attribute) {
}

func DrawSimpleBox(x, y, w, h int, fg, bg termbox.Attribute) {
	DrawBox(x, y, w, h, '+', '-', '|', fg, bg)
}

func DrawBox(startingX, startingY, w, h int, corner, top, side rune, fg, bg termbox.Attribute) {
	endX := w + startingX
	endY := h + startingY
	for x := startingX; x < endX; x++ {
		for y := startingY; y < endY; y++ {
			var char rune
			if (x == startingX || x == endX-1) && (y == startingY || y == endY-1) {
				char = corner
			} else if x == startingX || x == endX-1 {
				char = side
			} else if y == startingY || y == endY-1 {
				char = top
			} else {
				continue
			}
			termbox.SetCell(x, y, char, termbox.ColorBlue, termbox.ColorBlack)
		}
	}
}
