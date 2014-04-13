package particles

import (
	"eugor/dungeon"
	"github.com/nsf/termbox-go"
)

type Particle struct {
	Location     dungeon.Point
	velocity     dungeon.Point
	acceleration dungeon.Point
	lifespan     float64
}

func MakeParticle(position dungeon.Point) Particle {
	return Particle{
		Location:     position,
		velocity:     dungeon.MakePoint(0, 1),
		acceleration: dungeon.MakePoint(1, 1),
		lifespan:     10,
	}
}

func (p Particle) IsAlive() bool {
	return p.lifespan > 0
}

func (p Particle) Update() Particle {
	if p.IsAlive() {
		p.velocity = p.velocity.Plus(p.acceleration)
		p.Location = p.Location.Plus(p.velocity)
		p.lifespan -= 2
	}
	return p
}

func (p Particle) Draw() {
	if p.IsAlive() {
		termbox.SetCell(p.Location.X, p.Location.Y, 'p', termbox.ColorBlue, termbox.ColorGreen)
	}
}
