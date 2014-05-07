package dungeon

import (
	"github.com/nsf/termbox-go"
)

type Tile struct {
	Name         string
	Char         rune
	Walkable     bool
	Interactable bool
	SeeThrough   bool
	TransformsTo string
	Fg           termbox.Attribute
	Bg           termbox.Attribute
}

var Tiles = []Tile{
	Tile{Name: "floor", Char: ' ', Walkable: true, SeeThrough: true, Fg: termbox.ColorWhite, Bg: termbox.ColorBlack},
	Tile{Name: "wall", Char: 'X', Fg: termbox.ColorRed, Bg: termbox.ColorBlack},
	Tile{Name: "door", Char: 'D', Interactable: true, TransformsTo: "opendoor", Fg: termbox.ColorYellow, Bg: termbox.ColorBlack},
	Tile{Name: "opendoor", Char: '.', Interactable: true, Walkable: true, SeeThrough: true, TransformsTo: "door", Fg: termbox.ColorYellow, Bg: termbox.ColorBlack},
	Tile{Name: "greengrass", Char: '⁙', Walkable: true, Fg: termbox.ColorGreen, Bg: termbox.ColorBlack},
	Tile{Name: "bluegrass", Char: '⁙', Walkable: true, Fg: termbox.ColorBlue, Bg: termbox.ColorBlack},
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
	Width  int
	Height int
	Tiles  [][]uint16
	ViewX  int
	ViewY  int
}

func NewTileMap(width, height int) *TileMap {
	tiles := make([][]uint16, width)
	for i := range tiles {
		tiles[i] = make([]uint16, height)
	}
	return &TileMap{Width: width, Height: height, Tiles: tiles, ViewX: 0, ViewY: 0}
}

func (t *TileMap) DrawProjection(screenX, screenY int, positionX, positionY int) {
	if t.WithinRange(positionX, positionY) {
		value := t.Tiles[positionX][positionY]
		tile := Tiles[value]
		termbox.SetCell(screenX, screenY, tile.Char, tile.Fg, tile.Bg)
	}
}

func (t *TileMap) Draw() {
	width, height := termbox.Size()
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			pointX, pointY := x+t.ViewX, y+t.ViewY
			if t.WithinRange(pointX, pointY) {
				value := t.Tiles[x][y]
				tile := Tiles[value]
				termbox.SetCell(x, y, tile.Char, tile.Fg, tile.Bg)
			} else {
				var color termbox.Attribute
				if y%2 == 0 {
					color = termbox.ColorRed
				} else {
					color = termbox.ColorYellow
				}
				termbox.SetCell(x, y, '█', color, termbox.ColorBlack)
			}

		}
	}
}

func (t *TileMap) fill(x, y, width, height int, value uint16) {
	for startX := x; startX < x+width; startX++ {
		for startY := y; startY < y+height; startY++ {
			t.Tiles[startX][startY] = value
		}
	}
}

func (t *TileMap) FetchTile(x, y int) Tile {
	var tile Tile
	if t.WithinRange(x, y) {
		index := t.Tiles[x][y]
		tile = Tiles[index]
	} else {
		tile = Tiles[FindTile("wall")]
	}
	return tile
}

func (t *TileMap) WithinRange(x, y int) (within bool) {
	within = true
	if x < 0 || y < 0 {
		within = false
	} else if x >= len(t.Tiles) || y >= len(t.Tiles[x]) {
		within = false
	}
	return
}

func (t *TileMap) CanMoveTo(x, y int) bool {
	return t.FetchTile(x, y).Walkable
}

func (t *TileMap) CanInteractWith(x, y int) bool {
	return t.FetchTile(x, y).Interactable
}

func (t *TileMap) Interact(x, y int) {
	tile := t.FetchTile(x, y)
	replacement := FindTile(tile.TransformsTo)
	t.Tiles[x][y] = replacement
}

func (t *TileMap) AddRoom(x, y, width, height int) {
	var wall uint16 = FindTile("wall")
	t.fill(x, y, width, 1, wall)
	t.fill(x, y, 1, height, wall)
	t.fill(x+width-1, y, 1, height, wall)
	t.fill(x, y+height-1, width, 1, wall)
}

func (t *TileMap) AddDoor(x, y int) {
	var door uint16 = FindTile("door")
	t.Tiles[x][y] = door
}
