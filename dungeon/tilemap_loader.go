package dungeon

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func LoadTilemap(filename string) TileMap {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	width, height := extractDimensions(lines[0])
	tilemap := NewTileMap(width, height)
	return fillMap(tilemap, lines[1:])

}

func extractDimensions(line string) (int, int) {
	splitLine := strings.Split(line, "x")
	width, _ := strconv.ParseInt(splitLine[0], 10, 64)
	height, _ := strconv.ParseInt(splitLine[1], 10, 64)
	return int(width), int(height)
}

func fillMap(tilemap TileMap, lines []string) TileMap {
	y := 0
	for _, line := range lines {
		for x, letter := range line {
			tile, _ := strconv.ParseUint(string(letter), 10, 16)
			tilemap.Tiles[x][y] = uint16(tile)
		}
		y++
	}
	return tilemap
}
