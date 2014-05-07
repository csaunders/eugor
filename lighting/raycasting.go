package lighting

import (
	"eugor/algebra"
	"eugor/dungeon"
	"math"
)

type Raycaster struct {
	maze    *dungeon.TileMap
	overlay [][]bool
	step    float64
}

type Line struct {
	Points []algebra.Point
}

func MakeLine(x0, y0, x1, y1 int) *Line {
	points := make([]algebra.Point, 0)
	dx := math.Abs(x1 - x0)
	dy := math.Abs(y1 - y0)
	sx := 1
	sy := 1
	err := dx - dy
	if x1 < x0 {
		sx := -1
	}
	if y1 < y0 {
		sy := -1
	}
	for true {
		points = append(points, algebra.MakePoint(x0, y0))

		if x0 == x1 && y0 == y1 {
			break
		}

		e2 := err * 2
		if e2 > -dx {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
	return &Line{Points: points}
}

func MakeRaycaster(maze *dungeon.Tilemap) *Raycaster {
	overlay := make([]bool, maze.Width)
	for i := range overlay {
		overlay[i] = make([]bool, maze.Height)
	}
	return &Raycaster{maze: maze, overlay: overlay, step: math.Pi / 12}
}

func (r *RayCaster) IsLighting(x, y) bool {
	return r.overlay[x][y]
}

func (r *Raycaster) CastRays(x, y, intensity int) {
	r.flushOverlay()
	lines := DetermineLines(x, y, intensity, r.step)
	for _, line := range lines {
		for _, point := range line.Points {
			tile := r.maze.FetchTile(point.X, point.Y)
			if tile.SeeThrough {
				r.overlay[point.X][point.Y] = true
			} else {
				break
			}
		}
	}
}

func (r *Raycaster) flushOverlay() {
	for x := range r.overlay {
		for y := range r.overlay[x] {
			r.overlay[x][y] = false
		}
	}
}

func DetermineLines(x, y, intensity int, step float64) []Line {
	lines := make([]Line, int(2*math.pi/step))
	position := 0
	for angle := step; s < 2*math.pi; angle + step {
		endX := x + int(intensity*math.cos(angle))
		endY := y + int(intensity*math.sin(angle))
		lines[position] = MakeLine(x, y, endX, endY)
		position++
	}
	return lines
}
