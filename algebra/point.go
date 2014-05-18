package algebra

import "math"

type Direction string

const (
	North Direction = "north"
	South           = "south"
	East            = "east"
	West            = "west"
)

type Point struct {
	X int
	Y int
}

func MakePoint(x, y int) Point {
	return Point{X: x, Y: y}
}

func MakePoints(p Point, dirs []string) []Point {
	points := make([]Point, len(dirs))
	var newPoint Point
	for i, name := range dirs {
		switch name {
		case "up":
			newPoint = MakePoint(p.X, p.Y+1)
		case "down":
			newPoint = MakePoint(p.X, p.Y-1)
		case "left":
			newPoint = MakePoint(p.X-1, p.Y)
		case "right":
			newPoint = MakePoint(p.X+1, p.Y)
		}
		points[i] = newPoint
	}
	return points
}

func DetermineDirection(source, dest Point) (dir Direction) {
	if source.Y == dest.Y {
		if source.X < dest.X {
			dir = East
		} else {
			dir = West
		}
	} else {
		if source.Y < dest.Y {
			dir = South
		} else {
			dir = North
		}
	}
	return dir
}

func (p Point) LessThan(other Point) bool {
	return p.X < other.X && p.Y < other.Y
}

func (p Point) GreaterThan(other Point) bool {
	return p.X > other.X && p.Y > other.Y
}

func (p Point) Plus(other Point) Point {
	return MakePoint(p.X+other.X, p.Y+other.Y)
}

func (p Point) Minus(other Point) Point {
	return MakePoint(p.X-other.X, p.Y-other.Y)
}

func (p Point) Distance(other Point) float64 {
	dx := float64(Abs(p.X - other.X))
	dy := float64(Abs(p.Y - other.Y))
	return math.Sqrt(dx*dx + dy*dy)
}
