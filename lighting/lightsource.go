package lighting

type Lightsource interface {
	X() int
	Y() int
	Tick() Lightsource
	Intensity() int
}
