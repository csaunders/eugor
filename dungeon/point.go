package dungeon

type Point struct {
	X int
	Y int
}

func MakePoint(x, y int) Point {
	return Point{X: x, Y: y}
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
