package sprites

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func constTestAbility(successRate int) Ability {
	return Ability{Name: "test", Modifier: "test", SuccessRate: successRate, Constant: true}
}

func TestRollingADiceForAConstantAbility(t *testing.T) {
	ability := constTestAbility(60)
	success, value := ability.Roll()
	assert.Equal(t, true, success)
	assert.Equal(t, ability.SuccessRate, value)
}

func TestIssuingAChallengeToAWeakerAbility(t *testing.T) {
	attacker := constTestAbility(60)
	defender := constTestAbility(50)
	success := attacker.Challenge(defender)
	assert.Equal(t, true, success)
}

func TestIssuingAChallengeToAStrongerAbiity(t *testing.T) {
	attacker := constTestAbility(50)
	defender := constTestAbility(60)
	success := attacker.Challenge(defender)
	assert.Equal(t, false, success)
}

func TestIssuingAChallengeWhereTheAttackersRollFailed(t *testing.T) {
	attacker := Ability{Name: "fail", Modifier: "fail", SuccessRate: 0, Constant: false}
	defender := constTestAbility(1)
	success := attacker.Challenge(defender)
	assert.Equal(t, false, success)
}

func TestIssuingAChallengeWhereTheDefendersRollFailed(t *testing.T) {
	attacker := constTestAbility(1)
	defender := Ability{Name: "fail", Modifier: "fail", SuccessRate: 0, Constant: false}
	success := attacker.Challenge(defender)
	assert.Equal(t, true, success)
}
