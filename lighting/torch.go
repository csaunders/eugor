package lighting

import (
	"fmt"
	"math/rand"
)

type Torch struct {
	breathIn      bool
	baseIntensity int
	x             int
	y             int
}

func NewTorch(x, y int) Torch {
	return Torch{x: x, y: y, baseIntensity: 2, breathIn: false}
}

func (t Torch) X() int {
	return t.x
}

func (t Torch) Y() int {
	return t.y
}

func (t Torch) Intensity() (intensity int) {
	if t.breathIn {
		intensity = t.baseIntensity
	} else {
		intensity = t.baseIntensity + 1
	}
	return
}

func (t Torch) Tick() Lightsource {
	if rand.Int63n(100) <= 10 {
		t.breathIn = !t.breathIn
	}
	return t
}

func (t Torch) ToString() string {
	return fmt.Sprintf("(x: %d, y: %d, breathe: %v, intensity: %d)", t.x, t.y, t.breathIn, t.Intensity())
}
