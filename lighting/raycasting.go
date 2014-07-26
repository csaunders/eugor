package lighting

import (
	"eugor"
	"eugor/dungeon"
	"math"
)

type Raycaster struct {
	maze    *dungeon.TileMap
	overlay [][]bool
	step    float64
}

func MakeRaycaster(maze *dungeon.TileMap) *Raycaster {
	overlay := make([][]bool, maze.Width)
	for i := range overlay {
		overlay[i] = make([]bool, maze.Height)
	}
	return &Raycaster{maze: maze, overlay: overlay, step: math.Pi / 12}
}

func (r *Raycaster) IsLighting(x, y int) bool {
	return r.withinBounds(x, y) && r.overlay[x][y]
}

func (r *Raycaster) CastRays(x, y, intensity int) {
	r.FlushOverlay()
	r.calculateFieldOfView(x, y, intensity)
}

func (r *Raycaster) withinBounds(x, y int) bool {
	return x < r.maze.Width && x >= 0 && y < r.maze.Height && y >= 0
}

func (r *Raycaster) FlushOverlay() {
	for x := range r.overlay {
		for y := range r.overlay[x] {
			r.overlay[x][y] = false
		}
	}
}

func (r *Raycaster) calculateFieldOfView(x, y, intensity int) {
	r.overlay[x][y] = true
	r.sendRays(x, y, float64(intensity))
}

func (r *Raycaster) sendRays(fromX, fromY int, radius float64) {
	for theta := 0.0; theta <= 2*math.Pi; theta = theta + (math.Pi / 128) {
		toX := fromX + int(radius*math.Cos(theta))
		toY := fromY + int(radius*math.Sin(theta))
		r.DoLine(fromX, fromY, toX, toY)
	}
}

func (r *Raycaster) DoLine(x0, y0, x1, y1 int) {
	deltaX := eugor.Abs(x1 - x0)
	deltaY := eugor.Abs(y1 - y0)
	stepX := 1
	stepY := 1
	if x0 >= x1 {
		stepX = -1
	}
	if y0 >= y1 {
		stepY = -1
	}
	deltaErr := deltaX - deltaY
	for true {
		if !r.assignVisibility(x0, y0) {
			break
		}
		if x0 == x1 && y0 == y1 {
			break
		}

		deltaErr2 := 2 * deltaErr
		if deltaErr2 > -deltaY {
			deltaErr -= deltaY
			x0 += stepX
		}

		if deltaErr2 < deltaX {
			deltaErr += deltaX
			y0 += stepY
		}
	}
}

func (r *Raycaster) assignVisibility(x, y int) bool {
	if !r.withinBounds(x, y) {
		return false
	}
	tile := r.maze.FetchTile(x, y)
	r.overlay[x][y] = true
	return tile.SeeThrough
}
