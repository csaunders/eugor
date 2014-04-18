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
	writeLine(w, "[header]")
	writeLine(w, fmt.Sprintf("height=%d", data.Maze.Height))
	writeLine(w, fmt.Sprintf("width=%d", data.Maze.Width))
	appendEmptyLine(w)
}

func writePlayerInformation(w *bufio.Writer, data MapData) {
	writeLine(w, "[player]")
	writeLine(w, fmt.Sprintf("x=%d", data.PlayerStart.X))
	writeLine(w, fmt.Sprintf("y=%d", data.PlayerStart.Y))
	appendEmptyLine(w)
}

func writeLightSources(w *bufio.Writer, data MapData) {
	writeLine(w, "[lightsources]")
	for _, l := range data.MazeLights {
		writeLine(w, fmt.Sprintf("%s,%d,%d", l.Name(), l.X(), l.Y()))
	}
	appendEmptyLine(w)
}

func writeMapInformation(w *bufio.Writer, data MapData) {
	layer := TileMapToLayer(data.Maze)
	writeLayerInformation(w, layer)
}

func writeLayerInformation(w *bufio.Writer, layer LayerInformation) {
	writeLine(w, "[layer]")
	writeLine(w, fmt.Sprintf("type=%s", layer.Type))
	writeLine(w, "data=")
	for _, line := range layer.Data {
		writeLine(w, line)
	}
	appendEmptyLine(w)
}

func appendEmptyLine(w *bufio.Writer) {
	writeLine(w, "")
}

func writeLine(w *bufio.Writer, s string) {
	w.WriteString(s)
	w.WriteString("\n")
}
