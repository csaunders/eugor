package camera

import (
	"eugor/algebra"
	"eugor/dungeon"
	"github.com/nsf/termbox-go"
)

type Drawable interface {
	DrawProjection(screenX, screenY int, positionX, positionY int)
	X() int
	Y() int
}

func CameraDraw(field *dungeon.TileMap, focus Drawable, sprites []Drawable) (focusDrawPoint, fieldStartPoint algebra.Point, meta string) {
	var screenEnd = algebra.MakePoint(termbox.Size())
	if screenEnd.X < field.Width && screenEnd.Y < field.Height {
		focusDrawPoint, fieldStartPoint = determinePoints(focus.X(), focus.Y(), field.Width, field.Height)
	} else {
		focusDrawPoint = algebra.MakePoint(focus.X(), focus.Y())
		fieldStartPoint = algebra.MakePoint(0, 0)
	}

	for x := 0; x < screenEnd.X; x++ {
		for y := 0; y < screenEnd.Y; y++ {
			field.DrawProjection(x, y, fieldStartPoint.X+x, fieldStartPoint.Y+y)
		}
	}
	for _, sprite := range sprites {
		if IsOnScreen(sprite, fieldStartPoint) {
			x := sprite.X() - fieldStartPoint.X
			y := sprite.Y() - fieldStartPoint.Y
			sprite.DrawProjection(x, y, sprite.X(), sprite.Y())
		}
	}
	focus.DrawProjection(focusDrawPoint.X, focusDrawPoint.Y, focus.X(), focus.Y())
	return
}

func IsOnScreen(d Drawable, startPoint algebra.Point) bool {
	w, h := termbox.Size()
	x, y := d.X(), d.Y()
	if x > startPoint.X && y < (startPoint.X+w) {
		if x > startPoint.Y && y < (startPoint.Y+h) {
			return true
		}
	}
	return false
}

func determinePoints(x, y, fieldWidth, fieldHeight int) (focusDrawPoint, fieldStartPoint algebra.Point) {
	origin := algebra.MakePoint(0, 0)
	position := algebra.MakePoint(x, y)
	endField := algebra.MakePoint(fieldWidth, fieldHeight)
	w, h := termbox.Size()
	screenEnd := algebra.MakePoint(w, h)
	screenMid := algebra.MakePoint(screenEnd.X/2, screenEnd.Y/2)

	// TODO: Refactor the shit out of this, it maeks me sad
	switch {
	// Upper Left
	case position.Minus(screenMid).LessThan(origin):
		focusDrawPoint = position
		fieldStartPoint = origin
	// Upper Right
	case position.Y-screenMid.Y < 0 && endField.X-(position.X+screenMid.X) < 0:
		focusDrawPoint = algebra.MakePoint(position.X+screenEnd.X-endField.X, position.Y)
		fieldStartPoint = algebra.MakePoint(endField.X-screenEnd.X, 0)
	// Lower Left
	case position.X-screenMid.X < 0 && endField.Y-(position.Y+screenMid.Y) < 0:
		focusDrawPoint = algebra.MakePoint(position.X, position.Y+screenEnd.Y-endField.Y)
		fieldStartPoint = algebra.MakePoint(0, endField.Y-screenEnd.Y)
	// Lower Right
	case endField.Y-(position.Y+screenMid.Y) < 0 && endField.X-(position.X+screenMid.X) < 0:
		focusDrawPoint = algebra.MakePoint(position.X+screenEnd.X-endField.X, position.Y+screenEnd.Y-endField.Y)
		fieldStartPoint = algebra.MakePoint(endField.X-screenEnd.X, endField.Y-screenEnd.Y)
	// Upper Edge of Map
	case position.Y-screenMid.Y < 0:
		focusDrawPoint = algebra.MakePoint(screenMid.X, position.Y)
		fieldStartPoint = algebra.MakePoint(position.X-screenMid.X, 0)
	// Middle Left Edge of Map
	case position.X-screenMid.X < 0:
		focusDrawPoint = algebra.MakePoint(position.X, screenMid.Y)
		fieldStartPoint = algebra.MakePoint(0, position.Y-screenMid.Y)
	// Middle Right Edge of Map
	case endField.X-(position.X+screenMid.X) < 0:
		focusDrawPoint = algebra.MakePoint(position.X+screenEnd.X-endField.X, screenMid.Y)
		fieldStartPoint = algebra.MakePoint(endField.X-screenEnd.X, position.Y-screenMid.Y)
	// Bottom Edge of Map
	case endField.Y-(position.Y+screenMid.Y) < 0:
		focusDrawPoint = algebra.MakePoint(screenMid.X, position.Y+screenEnd.Y-endField.Y)
		fieldStartPoint = algebra.MakePoint(position.X-screenMid.X, endField.Y-screenEnd.Y)
	// Anywhere in the Middle of the Map
	default:
		focusDrawPoint = screenMid
		fieldStartPoint = algebra.MakePoint(position.X-screenMid.X, position.Y-screenMid.Y)
	}
	return
}
