package prng

import (
	"math/rand"
	"time"
)

var seed int64 = -1

func MakePrng() *rand.Rand {
	return rand.New(rand.NewSource(Seed()))
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
