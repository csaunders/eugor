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
}

func (hall Hallway) Dimensions() (x, y, w, h int) {
	x, y, w, h = 0, 0, hall.w, hall.h
	return
}

func (hall Hallway) Doors() map[Direction]Structure {
	return hall.rooms
}

func (hall Hallway) Attach(direction Direction, room Room) Hallway {
	return hall
}

func (hall Hallway) DetermineOffset(x, y int, direction Direction) (int, int) {
	return 0, 0
}
