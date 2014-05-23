package prng

import (
	"math/rand"
	"time"
)

func MakePrng() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
