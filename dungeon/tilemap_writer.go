package dungeon

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func SaveTilemap(data MapData, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	writeHeader(writer, data)
	writePlayerInformation(writer, data)
	writeLightSources(writer, data)
	writeMapInformation(writer, data)
}

func writeHeader(w *bufio.Writer, data MapData) {
	w.WriteString("[header]")
	w.WriteString(fmt.Sprintf("height=%d", data.Maze.Height))
	w.WriteString(fmt.Sprintf("width=%d", data.Maze.Width))
	appendEmptyLine(w)
}

func writePlayerInformation(w *bufio.Writer, data MapData) {
	w.WriteString("[player]")
	w.WriteString(fmt.Sprintf("x=%d", data.PlayerStart.X))
	w.WriteString(fmt.Sprintf("y=%d", data.PlayerStart.Y))
	appendEmptyLine(w)
}

func writeLightSources(w *bufio.Writer, data MapData) {
	w.WriteString("[lightsources]")
	for _, l := range data.MazeLights {
		w.WriteString(fmt.Sprintf("%s,%d,%d", l.Name(), l.X(), l.Y()))
	}
	appendEmptyLine(w)
}

func writeMapInformation(w *bufio.Writer, data MapData) {
	layer := TileMapToLayer(data.Maze)
	writeLayerInformation(w, layer)
}

func writeLayerInformation(w *bufio.Writer, layer LayerInformation) {
	w.WriteString("[layer]")
	w.WriteString(fmt.Sprintf("type=%s", layer.Type))
	w.WriteString("data=")
	for _, line := range layer.Data {
		w.WriteString(line)
	}
	appendEmptyLine(w)
}

func appendEmptyLine(w *bufio.Writer) {
	w.WriteString("")
}
