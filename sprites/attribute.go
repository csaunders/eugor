package sprites

import (
	"eugor/logger"
	"eugor/prng"
	"fmt"
)

type DieType int

const (
	d100 DieType = 100
	d20  DieType = 20
	d2   DieType = (2 * iota)
	d4
	d6
	d8
	d10
	d12
)

type CharacterAttributes struct {
	currentHp      int
	maxHp          int
	coreAttributes map[string]int
	abilities      map[string]Ability
}

type Ability struct {
	Name        string
	Modifier    string
	SuccessRate int
	Constant    bool
}

func (a Ability) Challenge(other Ability) (didSucceed bool) {
	didSucceed = false
	success, value := a.Roll()
	otherSuccess, otherValue := other.Roll()
	if success {
		didSucceed = true
		if otherSuccess {
			didSucceed = value > otherValue
		}
	}
	event := logger.Event{LogLevel: logger.Debug, Message: fmt.Sprintf("attacker: {h: %t, r: %d}, defender: {h: %t, r: %d}", success, value, otherSuccess, otherValue)}
	logger.GlobalLog.AppendEvent(event)
	return didSucceed
}

func (a Ability) Roll() (success bool, value int) {
	success = true

	if a.Constant {
		value = a.SuccessRate
	} else {
		value = roll(d100, 1)
		success = value > (100 - a.SuccessRate)
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
