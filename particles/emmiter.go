package particles

import (
	"eugor/dungeon"
)

type Emmiter struct {
	Origin    dungeon.Point
	particles []Particle
}

func MakeEmmiter(origin dungeon.Point) Emmiter {
	return Emmiter{
		Origin:    origin,
		particles: []Particle{},
	}
}

func (e Emmiter) AddParticle(p Particle) Emmiter {
	e.particles = append(e.particles, p)
	return e
}

func (e Emmiter) Update() Emmiter {
	remainingParticles := []Particle{}
	for _, p := range e.particles {
		if p.IsAlive() {
			p = p.Update()
			remainingParticles = append(remainingParticles, p)
		}
	}
	e.particles = remainingParticles
	e = e.AddParticle(MakeParticle(e.Origin))
	return e
}

func (e Emmiter) Draw() {
	for _, p := range e.particles {
		p.Draw()
	}
}
