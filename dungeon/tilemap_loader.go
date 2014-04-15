package dungeon

import (
	"bufio"
	"errors"
	"eugor/algebra"
	"eugor/lighting"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type MapData struct {
	Maze        TileMap
	PlayerStart algebra.Point
	MazeLights  []lighting.Lightsource
}

type LayerInformation struct {
	Type string
	Data []string
}

type TileMapParserState int

const (
	Unknown TileMapParserState = iota
	Header
	Player
	LightSources
	Layer
)

func LoadTilemap(filename string) TileMap {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	state := Unknown
	data := MapData{}
	for scanner.Scan() {
		var line string = ""
		switch state {
		case Header:
			data.Maze, line = prepareMaze(scanner)
		case Player:
			data.PlayerStart, line = extractPlayerDetails(scanner)
		case LightSources:
			data.MazeLights, line = extractLightSources(scanner)
		case Layer:
			var layer LayerInformation
			layer, line = extractLayerInformation(scanner)
			if len(layer.Data) > 0 {
				continue
			}
		}
		if len(line) == 0 {
			line = scanner.Text()
		}
		state = determineState(line)
	}

	return data.Maze
}

func prepareMaze(scanner *bufio.Scanner) (TileMap, string) {
	// state := Header
	var width, height int = 0, 0
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		if determineState(line) == Unknown {
			break
		}
		splitLine := strings.Split(line, "=")
		if len(splitLine) > 2 {
			log.Fatal(errors.New(fmt.Sprintf("Invalid header information for %s", line)))
		}
	}
	return NewTileMap(width, height), ""
}

func extractPlayerDetails(scanner *bufio.Scanner) (algebra.Point, string) {
	return algebra.MakePoint(0, 0), ""
}

func extractLightSources(scanner *bufio.Scanner) ([]lighting.Lightsource, string) {
	return []lighting.Lightsource{}, ""
}

func extractLayerInformation(scanner *bufio.Scanner) (LayerInformation, string) {
	return LayerInformation{}, ""
}

func determineState(line string) TileMapParserState {
	return Unknown
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
