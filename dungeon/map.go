package dungeon

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
}

type Map struct {
	root      Structure
	viewportX int
	viewportY int
}

func NewMap() Map {
	m := Map{viewportX: 0, viewportY: 0}
	r := NewRoom(20, 20, 10, 10)
	r = r.Attach(North, NewHallway(3, 10))
	r = r.Attach(South, NewHallway(3, 10))
	r = r.Attach(East, NewHallway(15, 3))
	r = r.Attach(West, NewHallway(15, 3))
	m.root = r
	return m
}

func (m Map) Draw() {
	m.root.Draw(m.viewportX, m.viewportY)
}
