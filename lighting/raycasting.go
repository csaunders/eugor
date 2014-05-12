package lighting

import (
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
	return r.overlay[x][y]
}

func (r *Raycaster) CastRays(x, y, intensity int) {
	r.flushOverlay()
	r.calculateFieldOfView(x, y, intensity)
}

func (r *Raycaster) flushOverlay() {
	for x := range r.overlay {
		for y := range r.overlay[x] {
			r.overlay[x][y] = false
		}
	}
}

var OCTANT_CALCULATIONS map[int][]int = map[int][]int{
	0:  {1, 1, 1, 0},
	1:  {1, -1, 1, 0},
	2:  {-1, 1, 1, 0},
	3:  {-1, -1, 1, 0},
	4:  {1, 1, 0, 1},
	5:  {1, -1, 0, 1},
	6:  {-1, 1, 0, 1},
	7:  {-1, -1, 0, 1},
	8:  {1, 0, 1, 0},
	9:  {0, 1, 1, 0},
	10: {-1, 0, 0, 1},
	11: {0, -1, 0, 1},
}

func (r *Raycaster) calculateFieldOfView(x, y, intensity int) {
	for _, values := range OCTANT_CALCULATIONS {
		sx, sy, dx, dy := values[0], values[1], values[2], values[3]
		r.doOctant(x, y, intensity, sx, sy, dx, dy)
	}
}

// Algorithm from Rogue Basin
// http://www.roguebasin.com/index.php?title=Ray-Tracing_Field-Of-View_Demo
func (r *Raycaster) doOctant(x, y, radius, sx, sy, dx, dy int) {
	for i := 0; i != radius; i++ {
		var lastTile *dungeon.Tile
		var lastAdjacentTile *dungeon.Tile
		for j := 0; j != radius; j++ {
			tileX := x + (sx * i)
			tileY := y + (sy * j)
			if tileX >= r.maze.Width || tileX < 0 || tileY >= r.maze.Height || tileY < 0 {
				break
			}
			tile := r.maze.FetchTile(tileX, tileY)

			adjacentTile := r.maze.FetchTile(tileX-(sx*dx), tileY-(sy*dy))
			if lastTile != nil {
				if lastTile.Name != "wall" {
					r.overlay[tileX][tileY] = true
				} else {
					if tileX <= 0 {
						break
					}

					tileIsWall := tile != nil && tile.Name == "wall"
					adjacentTileIsClear := adjacentTile != nil && adjacentTile.Name != "wall"
					lastAdjacentTileIsClear := lastAdjacentTile != nil && lastAdjacentTile.Name != "wall"

					if tileIsWall && adjacentTileIsClear && lastAdjacentTileIsClear {
						r.overlay[tileX][tileY] = true
					} else {
						break
					}
				}
			}
			lastTile = tile
			lastAdjacentTile = adjacentTile
		}
	}
}
