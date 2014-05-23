package sprites

import (
	"eugor/prng"
)

type DieType int

const (
	d20 DieType = 20
	d2  DieType = (2 * iota)
	d4
	d6
	d8
	d10
	d12
)

type Attribute struct {
	currentHp    int
	maxHp        int
	strength     int
	dexterity    int
	constitution int
}

type AttackAttribute struct {
	numDice        int
	dieType        DieType
	damageModifier int
	hitModifier    int
}

type DefenseAttribute struct {
	armorClass int
}

func Attack(attack AttackAttribute, defense DefenseAttribute) (damage int, didHit bool) {
	damage = 0
	didHit = false
	if roll(d20, 1)+attack.hitModifier > defense.armorClass {
		damage = roll(attack.dieType, attack.numDice) + attack.damageModifier
		didHit = true
	}
	return
}

func roll(dieType DieType, number int) int {
	rng := prng.MakePrng()
	sum := 0
	for i := 0; i < number; i++ {
		sum += rng.Intn(int(dieType)) + 1
	}
	return sum
}
