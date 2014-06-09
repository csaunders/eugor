package prng

import (
	"math/rand"
	"time"
)

var seed int64 = -1
var prng *rand.Rand

func MakePrng() *rand.Rand {
	if prng == nil {
		prng = rand.New(rand.NewSource(Seed()))
	}
	return prng
}

func Seed() int64 {
	if seed <= 0 {
		seed = time.Now().UnixNano()
	}
	return seed
}

// Don't actually use this -- it's for replicating
// PRNG issues
func SetSeed(newSeed int64) {
	seed = newSeed
}
