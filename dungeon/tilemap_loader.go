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

type TileMapParserState int

const (
	Unknown TileMapParserState = iota
	Continue
	Header
	Player
	LightSources
	Layer
)

func LoadTilemap(filename string) MapData {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	state := Unknown
	data := MapData{}
	var line string = ""
	for scanner.Scan() {
		switch state {
		case Header:
			data.Maze, line = prepareMaze(scanner)
			fmt.Println("Maze Overview has been Extracted")
		case Player:
			data.PlayerStart, line = extractPlayerDetails(scanner)
			fmt.Println("Player Overview has been Extracted")
		case LightSources:
			data.MazeLights, line = extractLightSources(scanner)
			fmt.Println("Light Sources have been Extracted")
		case Layer:
			var layer LayerInformation
			layer, line = extractLayerInformation(scanner)
			fmt.Println("Layer details have been Extracted")
			if layer.Type == "map" && len(layer.Data) > 0 {
				fmt.Println("Layer is a map, filling Maze with data")
				fillMap(data.Maze, layer.Data)
			}
		}
		if len(line) == 0 {
			line = scanner.Text()
		}
		state = determineState(line)
	}

	return data
}

func prepareMaze(scanner *bufio.Scanner) (*TileMap, string) {
	var width, height int = 0, 0
	for true {
		line := scanner.Text()
		if determineState(line) == Unknown {
			break
		}
		splitLine := strings.Split(line, "=")
		if len(splitLine) > 2 {
			log.Fatal(errors.New(fmt.Sprintf("Invalid header information for %s", line)))
		}
		value, _ := strconv.ParseInt(splitLine[1], 10, 64)
		switch splitLine[0] {
		case "width":
			width = int(value)
		case "height":
			height = int(value)
		}
		scanner.Scan()
	}
	return NewTileMap(width, height), ""
}

func extractPlayerDetails(scanner *bufio.Scanner) (algebra.Point, string) {
	var x, y int = 0, 0
	for true {
		line := scanner.Text()
		if determineState(line) == Unknown {
			break
		}
		splitLine := strings.Split(line, "=")
		if len(splitLine) > 2 {
			log.Fatal(errors.New(fmt.Sprintf("Invalid player information for %s", line)))
		}
		value, _ := strconv.ParseInt(splitLine[1], 10, 64)
		switch splitLine[0] {
		case "x":
			x = int(value)
		case "y":
			y = int(value)
		}
		scanner.Scan()
	}
	return algebra.MakePoint(x, y), ""
}

func extractLightSources(scanner *bufio.Scanner) ([]lighting.Lightsource, string) {
	lights := []lighting.Lightsource{}
	for true {
		line := scanner.Text()
		if determineState(line) == Unknown {
			break
		}
		splitLine := strings.Split(line, ",")
		if len(splitLine) > 3 {
			log.Fatal(errors.New(fmt.Sprintf("Invalid light source information for %s", line)))
		}
		switch splitLine[0] {
		case "torch":
			x, _ := strconv.ParseInt(splitLine[1], 10, 64)
			y, _ := strconv.ParseInt(splitLine[2], 10, 64)
			torch := lighting.NewTorch(int(x), int(y))
			lights = append(lights, torch)
		}

		scanner.Scan()
	}
	return lights, ""
}

func extractLayerInformation(scanner *bufio.Scanner) (LayerInformation, string) {
	layer := LayerInformation{}
	scanning := true
	for scanning {
		line := scanner.Text()
		if determineState(line) == Unknown {
			scanning = false
			break
		}
		splitLine := strings.Split(line, "=")
		if len(splitLine) < 1 {
			log.Fatal(errors.New(fmt.Sprintf("Invalid layer information for %s", line)))
		}
		switch splitLine[0] {
		case "type":
			layer.Type = splitLine[1]
		case "data":
			layer.Data = extractMapData(scanner)
			scanning = false
		}
		if scanning {
			scanner.Scan()
		}
	}
	fmt.Println("returning layer")
	return layer, ""
}

func determineState(line string) TileMapParserState {
	switch line {
	case "[header]":
		return Header
	case "[player]":
		return Player
	case "[lightsources]":
		return LightSources
	case "[layer]":
		return Layer
	case "":
		return Unknown
	default:
		return Continue
	}
}

func extractDimensions(line string) (int, int) {
	splitLine := strings.Split(line, "x")
	width, _ := strconv.ParseInt(splitLine[0], 10, 64)
	height, _ := strconv.ParseInt(splitLine[1], 10, 64)
	return int(width), int(height)
}

func extractMapData(scanner *bufio.Scanner) []string {
	result := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if determineState(line) == Unknown {
			break
		}
		result = append(result, line)
	}
	return result
}

func fillMap(tilemap *TileMap, lines []string) {
	y := 0
	for _, line := range lines {
		entries := strings.Split(line, ",")
		for x, letter := range entries {
			tile, _ := strconv.ParseUint(string(letter), 10, 16)
			tilemap.Tiles[x][y] = uint16(tile)
		}
		y++
	}
}
