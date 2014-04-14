package particles

import (
	"eugor/algebra"
)

type Emmiter struct {
	Origin    algebra.Point
	Particles []Particle
	birthRate int
}

func MakeEmmiter(origin algebra.Point, birthRate int) Emmiter {
	return Emmiter{
		Origin:    origin,
		Particles: []Particle{},
		birthRate: birthRate,
	}
}

func (e Emmiter) AddParticle(p Particle) Emmiter {
	e.Particles = append(e.Particles, p)
	return e
}

func (e Emmiter) spawnMoreParticles() Emmiter {
	for i := 0; i < e.birthRate; i++ {
		newParticle := MakeParticle(e.Origin)
		e.Particles = append(e.Particles, newParticle)
	}
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
	e = e.spawnMoreParticles()
	return e
}

func (e Emmiter) Draw() {
	for _, p := range e.Particles {
		p.Draw()
	}
}
