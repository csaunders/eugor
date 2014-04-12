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
	origin := dungeon.MakePoint(0, 0)
	position := dungeon.MakePoint(focus.X(), focus.Y())
	w, h := termbox.Size()
	screenEnd := dungeon.MakePoint(w, h)
	screenMid := dungeon.MakePoint(screenEnd.X/2, screenEnd.Y/2)
	endField := dungeon.MakePoint(field.Width, field.Height)

	// TODO: Refactor the shit out of this, it maeks me sad
	switch {
	// Upper Left
	case position.Minus(screenMid).LessThan(origin):
		focusDrawPoint = position
		fieldStartPoint = origin
	// Upper Right
	case position.Y-screenMid.Y < 0 && endField.X-(position.X+screenMid.X) < 0:
		focusDrawPoint = dungeon.MakePoint(position.X+screenEnd.X-endField.X, position.Y)
		fieldStartPoint = dungeon.MakePoint(endField.X-screenEnd.X, 0)
	// Lower Left
	case position.X-screenMid.X < 0 && endField.Y-(position.Y+screenMid.Y) < 0:
		focusDrawPoint = dungeon.MakePoint(position.X, position.Y+screenEnd.Y-endField.Y)
		fieldStartPoint = dungeon.MakePoint(0, endField.Y-screenEnd.Y)
	// Lower Right
	case endField.Y-(position.Y+screenMid.Y) < 0 && endField.X-(position.X+screenMid.X) < 0:
		focusDrawPoint = dungeon.MakePoint(position.X+screenEnd.X-endField.X, position.Y+screenEnd.Y-endField.Y)
		fieldStartPoint = dungeon.MakePoint(endField.X-screenEnd.X, endField.Y-screenEnd.Y)
	// Upper Edge of Map
	case position.Y-screenMid.Y < 0:
		focusDrawPoint = dungeon.MakePoint(screenMid.X, position.Y)
		fieldStartPoint = dungeon.MakePoint(position.X-screenMid.X, 0)
	// Middle Left Edge of Map
	case position.X-screenMid.X < 0:
		focusDrawPoint = dungeon.MakePoint(position.X, screenMid.Y)
		fieldStartPoint = dungeon.MakePoint(0, position.Y-screenMid.Y)
	// Middle Right Edge of Map
	case endField.X-(position.X+screenMid.X) < 0:
		focusDrawPoint = dungeon.MakePoint(position.X+screenEnd.X-endField.X, screenMid.Y)
		fieldStartPoint = dungeon.MakePoint(endField.X-screenEnd.X, position.Y-screenMid.Y)
	// Bottom Edge of Map
	case endField.Y-(position.Y+screenMid.Y) < 0:
		focusDrawPoint = dungeon.MakePoint(screenMid.X, position.Y+screenEnd.Y-endField.Y)
		fieldStartPoint = dungeon.MakePoint(position.X-screenMid.X, endField.Y-screenEnd.Y)
	// Anywhere in the Middle of the Map
	default:
		focusDrawPoint = screenMid
		fieldStartPoint = dungeon.MakePoint(position.X-screenMid.X, position.Y-screenMid.Y)
	}

	for x := 0; x < screenEnd.X; x++ {
		for y := 0; y < screenEnd.Y; y++ {
			field.DrawProjection(x, y, fieldStartPoint.X+x, fieldStartPoint.Y+y)
		}
	}
	focus.DrawProjection(focusDrawPoint.X, focusDrawPoint.Y, position.X, position.Y)
	return
}
