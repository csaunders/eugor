package dungeon

import (
	"eugor/termboxext"
	"github.com/nsf/termbox-go"
)

type Room struct {
	x, y, w, h int
	doors      map[Direction]Structure
	Color      termbox.Attribute
}

func NewRoom(x, y, w, h int) Room {
	return Room{x: x, y: y, w: w, h: h, doors: make(map[Direction]Structure), Color: termbox.ColorRed}
}

func (r Room) Draw(x, y int) {
	termboxext.DrawSimpleBox(x+r.x, y+r.y, r.w, r.h, r.Color, termbox.ColorBlack)
	for direction, hallway := range r.doors {
		offsetX, offsetY := r.DetermineOffset(x, y, direction)
		hallway.Draw(offsetX, offsetY)
	}
}

func (r Room) Dimensions() (x, y, w, h int) {
	return r.x, r.y, r.w, r.h
}

func (r Room) Doors() map[Direction]Structure {
	return r.doors
}

func (r Room) DetermineOffset(x, y int, direction Direction) (int, int) {
	var offX, offY int
	switch direction {
	case North:
		offY = y + r.y - (r.h - 1)
		offX = x + r.x + (r.w / 2)
	case South:
		offY = y + r.y + r.h - 1
		offX = x + r.x + (r.h / 2)
	case East:
		offX = x + r.x + r.w - 1
		offY = y + r.y + (r.h / 2)
	case West:
		offX = x + ((r.x - r.w) / 2) + 1
		offY = y + r.y + (r.h / 2)
	}
	return offX, offY
}

func (r Room) IsWithinBounds(offsetX, offsetY, x, y int) bool {
	actualX, actualY := r.x+offsetX, r.y+offsetY
	withinX := x > actualX && x < actualX+r.w
	withinY := y > actualY && y < actualY+r.h
	if withinX && withinY {
		return true
	}
	for direction, hall := range r.doors {
		hallOffX, hallOffY := r.DetermineOffset(offsetX, offsetX, direction)
		if hall.IsWithinBounds(hallOffX, hallOffY, x, y) {
			return true
		}
	}
	return false
}

func (r Room) Attach(direction Direction, hallway Hallway) Room {
	r.doors[direction] = hallway
	return r
}
