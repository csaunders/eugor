package dungeon

import (
	"eugor/termboxext"
	"github.com/nsf/termbox-go"
)

type Hallway struct {
	w, h  int
	rooms map[Direction]Structure
}

func NewHallway(w, h int) Hallway {
	return Hallway{w: w, h: h, rooms: make(map[Direction]Structure)}
}

func (hall Hallway) Draw(x, y int) {
	termboxext.DrawSimpleBox(x, y, hall.w, hall.h, termbox.ColorRed, termbox.ColorBlack)
	for direction, room := range hall.rooms {
		switch direction {
		case North:
			room.Draw(x, y-hall.h+1)
		case South:
			room.Draw(x, y+hall.h-1)
		case East:
			room.Draw(x+hall.w-1, y)
		case West:
			room.Draw(x-hall.w+1, y)
		}
	}
}

func (hall Hallway) Dimensions() (x, y, w, h int) {
	x, y, w, h = 0, 0, hall.w, hall.h
	return
}

func (hall Hallway) Doors() map[Direction]Structure {
	return hall.rooms
}

func (hall Hallway) IsWithinBounds(offsetX, offsetY, x, y int) bool {
	withinX := x > offsetX && x < offsetX+hall.w
	withinY := y > offsetY && y < offsetY+hall.h
	if withinX && withinY {
		return true
	}
	for direction, room := range hall.rooms {
		roomOffX, roomOffY := hall.DetermineOffset(offsetX, offsetX, direction)
		if room.IsWithinBounds(roomOffX, roomOffY, x, y) {
			return true
		}
	}
	return false
}

func (hall Hallway) Attach(direction Direction, room Room) Hallway {
	hall.rooms[direction] = room
	return hall
}

func (hall Hallway) DetermineOffset(x, y int, direction Direction) (int, int) {
	return 0, 0
}
