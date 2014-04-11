package camera

import (
	"eugor/dungeon"
	"github.com/nsf/termbox-go"
)

type Drawable interface {
	DrawProjection(screenX, screenY int, positionX, positionY int)
	X() int
	Y() int
}

func CameraDraw(field dungeon.TileMap, focus Drawable) (focusDrawPoint, fieldStartPoint dungeon.Point, meta string) {
	position := dungeon.MakePoint(focus.X(), focus.Y())
	w, h := termbox.Size()
	screenEnd := dungeon.MakePoint(w, h)
	screenMid := dungeon.MakePoint(screenEnd.X/2, screenEnd.Y/2)
	endField := dungeon.MakePoint(field.Width, field.Height)
	if position.LessThan(screenMid) {
		meta = "Less than midscreen"
		focusDrawPoint = position
		fieldStartPoint = dungeon.MakePoint(0, 0)
	} else if position.GreaterThan(endField.Minus(screenMid)) {
		meta = "Greater than field size"
		focusDrawPoint = position.Minus(endField)
		fieldStartPoint = endField.Minus(screenEnd)
	} else {
		meta = "Somewhere in between"
		focusDrawPoint = screenMid
		if position.X < screenMid.X {
			focusDrawPoint.X = position.X
		}
		if position.Y < screenMid.Y {
			focusDrawPoint.Y = position.Y
		}
		fieldStartPoint = position.Minus(screenMid)
	}
	if fieldStartPoint.X < 0 {
		fieldStartPoint.X = 0
	}
	if fieldStartPoint.Y < 0 {
		fieldStartPoint.Y = 0
	}
	for x := 0; x < screenEnd.X; x++ {
		for y := 0; y < screenEnd.Y; y++ {
			field.DrawProjection(x, y, fieldStartPoint.X+x, fieldStartPoint.Y+y)
		}
	}
	focus.DrawProjection(focusDrawPoint.X, focusDrawPoint.Y, position.X, position.Y)
	return
}
