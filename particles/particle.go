package particles

import (
	"eugor/dungeon"
	"github.com/nsf/termbox-go"
	"math/rand"
	"time"
)

type Particle struct {
	Location dungeon.Point
	velocity dungeon.Point
	lifespan float64
}

func MakeParticle(position dungeon.Point) Particle {
	prng := rand.New(rand.NewSource(time.Now().UnixNano()))
	vx := prng.Intn(2)
	vy := prng.Intn(2)
	if prng.Intn(2) == 0 {
		vx = -vx
	}
	if prng.Intn(2) == 0 {
		vy = -vy
	}

	return Particle{
		Location: position,
		velocity: dungeon.MakePoint(vx, vy),
		lifespan: 5,
	}
}

func (p Particle) IsAlive() bool {
	return p.lifespan > 0
}

func (p Particle) IsDead() bool {
	return !p.IsAlive()
}

func (p Particle) Update() Particle {
	if p.IsAlive() {
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
