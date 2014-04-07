package dungeon

import (
	"eugor/sprites"
	"github.com/nsf/termbox-go"
)

type Direction string

const (
	North Direction = "north"
	South Direction = "south"
	East  Direction = "east"
	West  Direction = "west"
)

type Structure interface {
	Draw(x, y int)
	Dimensions() (x, y, w, h int)
	Doors() map[Direction]Structure
	DetermineOffset(x, y int, direction Direction) (int, int)
	IsWithinBounds(offsetX, offsetY, x, y int) bool
}

type Map struct {
	root      Structure
	viewportX int
	viewportY int
}

func NewMap() Map {
	m := Map{viewportX: 0, viewportY: 0}
	r := NewRoom(20, 20, 10, 10)
	hN := NewHallway(3, 10)
	hN = hN.Attach(North, NewRoom(0, 0, 15, 10))
	r = r.Attach(North, hN)
	r = r.Attach(South, NewHallway(3, 10))
	hE := NewHallway(15, 3)
	hE = hE.Attach(East, NewRoom(0, 0, 40, 25))
	r = r.Attach(East, hE)
	r = r.Attach(West, NewHallway(15, 3))
	m.root = r
	return m
}

func (m Map) Move(k termbox.Key) Map {
	switch {
	case k == termbox.KeyArrowUp:
		m.viewportY -= 1
	case k == termbox.KeyArrowDown:
		m.viewportY += 1
	case k == termbox.KeyArrowLeft:
		m.viewportX -= 1
	case k == termbox.KeyArrowRight:
		m.viewportX += 1
	}
	return m
}

func (m Map) IsWithinBounds(c sprites.Character) bool {
	return m.root.IsWithinBounds(m.viewportX, m.viewportY, c.X-m.viewportX, c.Y-m.viewportY)
}

func (m Map) Draw() {
	m.root.Draw(m.viewportX, m.viewportY)
}
