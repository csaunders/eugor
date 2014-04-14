package particles

import (
	"eugor/algebra"
)

type Emmiter struct {
	Origin    algebra.Point
	Particles []Particle
}

func MakeEmmiter(origin algebra.Point) Emmiter {
	return Emmiter{
		Origin:    origin,
		Particles: []Particle{},
	}
}

func (e Emmiter) AddParticle(p Particle) Emmiter {
	e.Particles = append(e.Particles, p)
	return e
}

func (e Emmiter) Update() Emmiter {
	remainingParticles := []Particle{}
	for _, p := range e.Particles {
		if p.IsAlive() {
			p = p.Update()
			remainingParticles = append(remainingParticles, p)
		}
	}
	e.Particles = remainingParticles
	return e
}

func (e Emmiter) Draw() {
	for _, p := range e.Particles {
		p.Draw()
	}
}
