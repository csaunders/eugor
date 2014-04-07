package dungeon

import (
	"github.com/nsf/termbox-go"
)

type Tile struct {
	Name         string
	Char         rune
	Walkable     bool
	Interactable bool
	TransformsTo string
	Fg           termbox.Attribute
	Bg           termbox.Attribute
}

var Tiles = []Tile{
	Tile{Name: "floor", Char: ' ', Walkable: true, Fg: termbox.ColorWhite, Bg: termbox.ColorBlack},
	Tile{Name: "wall", Char: 'X', Fg: termbox.ColorRed, Bg: termbox.ColorBlack},
	Tile{Name: "door", Char: 'D', Interactable: true, TransformsTo: "opendoor", Fg: termbox.ColorYellow, Bg: termbox.ColorBlack},
	Tile{Name: "opendoor", Char: '.', Interactable: true, Walkable: true, TransformsTo: "door", Fg: termbox.ColorYellow, Bg: termbox.ColorBlack},
}

func FindTile(name string) uint16 {
	for index, tile := range Tiles {
		if tile.Name == name {
			return uint16(index)
		}
	}
	return uint16(0)
}

type TileMap struct {
	Tiles [][]uint16
	ViewX int
	ViewY int
}

func NewTileMap(width, height int) TileMap {
	tiles := make([][]uint16, width)
	for i := range tiles {
		tiles[i] = make([]uint16, height)
	}
	tileMap := TileMap{Tiles: tiles, ViewX: 0, ViewY: 0}
	return tileMap
}

func (t TileMap) Draw() {
	for x := range t.Tiles {
		for y := range t.Tiles[x] {
			value := t.Tiles[x][y]
			tile := Tiles[value]
			termbox.SetCell(x, y, tile.Char, tile.Fg, tile.Bg)
		}
	}
}

func (t TileMap) fill(x, y, width, height int, value uint16) TileMap {
	for startX := x; startX < x+width; startX++ {
		for startY := y; startY < y+height; startY++ {
			t.Tiles[startX][startY] = value
		}
	}
	return t
}

func (t TileMap) fetchTile(x, y int) Tile {
	index := t.Tiles[x][y]
	return Tiles[index]
}

func (t TileMap) CanMoveTo(x, y int) bool {
	return t.fetchTile(x, y).Walkable
}

func (t TileMap) CanInteractWith(x, y int) bool {
	return t.fetchTile(x, y).Interactable
}

func (t TileMap) Interact(x, y int) TileMap {
	tile := t.fetchTile(x, y)
	replacement := FindTile(tile.TransformsTo)
	t.Tiles[x][y] = replacement
	return t
}

func (t TileMap) AddRoom(x, y, width, height int) TileMap {
	var wall uint16 = FindTile("wall")
	t = t.fill(x, y, width, 1, wall)
	t = t.fill(x, y, 1, height, wall)
	t = t.fill(x+width-1, y, 1, height, wall)
	t = t.fill(x, y+height-1, width, 1, wall)
	return t
}

func (t TileMap) AddDoor(x, y int) TileMap {
	var door uint16 = FindTile("door")
	t.Tiles[x][y] = door
	return t
}
