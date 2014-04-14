package lighting

import (
	"eugor/algebra"
	"eugor/particles"
	"fmt"
)

type Torch struct {
	emmiter       particles.Emmiter
	breathIn      bool
	baseIntensity int
	x             int
	y             int
}

func NewTorch(x, y int) Torch {
	point := algebra.MakePoint(x, y)
	emmiter := particles.MakeEmmiter(point)
	return Torch{emmiter: emmiter, x: x, y: y}
}

func (t Torch) X() int {
	return t.x
}

func (t Torch) Y() int {
	return t.y
}

func (t Torch) Intensity() (intensity int) {
	intensity = 0
	return
}

func (t Torch) IsLighting(x, y int) bool {
	for _, particle := range t.emmiter.Particles {
		loc := particle.Location
		if loc.X == x && loc.Y == y {
			return true
		}
	}
	return false
}

func (t Torch) Tick() Lightsource {
	t.emmiter = t.emmiter.Update()
	t.emmiter = t.emmiter.AddParticle(particles.MakeParticle(t.emmiter.Origin))
	t.emmiter = t.emmiter.AddParticle(particles.MakeParticle(t.emmiter.Origin))
	t.emmiter = t.emmiter.AddParticle(particles.MakeParticle(t.emmiter.Origin))
	t.emmiter = t.emmiter.AddParticle(particles.MakeParticle(t.emmiter.Origin))
	return t
}

func (t Torch) ToString() string {
	return fmt.Sprintf("(x: %d, y: %d, breathe: %v, intensity: %d)", t.x, t.y, t.breathIn, t.Intensity())
}
