package dungeon

import (
	"math/rand"
	"time"
)

var generatorPrng *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateTileMap(width, height int) TileMap {
	maze := NewTileMap(width, height)
	tiles := maze.Tiles
	fillTileMap(FindTile("wall"), tiles)
	digStartingRoom(maze, tiles)

	maze.Tiles = tiles
	return maze
}

func fillTileMap(value uint16, tiles [][]uint16) {
	for y, _ := range tiles {
		for x, _ := range tiles[y] {
			tiles[y][x] = value
		}
	}
}

func digStartingRoom(maze TileMap, tiles [][]uint16) {
	originX := maze.Width / 2
	originY := maze.Height / 2
	width := 14 + generatorPrng.Intn(5)
	height := 8 + generatorPrng.Intn(5)
	digRoom(tiles, originX, originY, width, height)
}

func digRoom(tiles [][]uint16, oX, oY, width, height int) {
	startX := oX - (width / 2)
	startY := oY - (height / 2)
	floor := FindTile("floor")
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			tiles[x+startX][y+startY] = floor
		}
	}
}
