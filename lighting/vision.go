package lighting

import (
	"math"
)

type Vision struct {
	intensity int
	x         int
	y         int
}

func NewVision(x, y, intensity int) Vision {
	return Vision{x: x, y: y, intensity: intensity}
}

func (v Vision) X() int {
	return v.x
}

func (v Vision) Y() int {
	return v.y
}

func (v Vision) Intensity() int {
	return v.intensity
}

func (v Vision) IsLighting(x, y int) bool {
	return math.Abs(float64(v.x-x)) <= float64(v.intensity) && math.Abs(float64(v.y-y)) <= float64(v.intensity)
}

func (v Vision) Tick() Lightsource {
	return v
}

func (v Vision) Projection() Projection {
	return Static
}

func (v Vision) ToString() string {
	return ""
}
