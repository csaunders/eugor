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

func MakeLine(x0, y0, x1, y1 float64) Line {
	points := []algebra.Point{}
	dx := math.Abs(x1 - x0)
	dy := math.Abs(y1 - y0)
	sx := 1.0
	sy := 1.0
	err := dx - dy
	if x1 < x0 {
		sx = -1.0
	}
	if y1 < y0 {
		sy = -1.0
	}
	max := 1000
	i := 0
	for true {
		points = append(points, algebra.MakePoint(int(x0), int(y0)))

		if i > max {
			break
		}
		i++

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
	return Line{Points: points}
}

func MakeRaycaster(maze *dungeon.TileMap) *Raycaster {
	overlay := make([][]bool, maze.Width)
	for i := range overlay {
		overlay[i] = make([]bool, maze.Height)
	}
	return &Raycaster{maze: maze, overlay: overlay, step: math.Pi / 12}
}

func (r *Raycaster) IsLighting(x, y int) bool {
	return r.overlay[x][y]
}

func (r *Raycaster) CastRays(x, y, intensity int) {
	r.flushOverlay()
	lines := DetermineLines(float64(x), float64(y), float64(intensity), r.step)
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

func DetermineLines(x, y, intensity, step float64) []Line {
	lines := []Line{}
	for angle := step; step <= 2*math.Pi; angle = angle + step {
		endX := x + (intensity * math.Cos(angle))
		endY := y + (intensity * math.Sin(angle))
		lines = append(lines, MakeLine(x, y, endX, endY))
	}
	return lines
}
