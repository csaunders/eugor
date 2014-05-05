package lighting

type Projection int

const (
	Static Projection = iota
	Relative
)

type Lightsource interface {
	X() int
	Y() int
	Tick()
	Intensity() int
	Projection() Projection
	ToString() string
	Name() string
	IsLighting(x, y int) bool
}
