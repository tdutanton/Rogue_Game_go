package common

import (
	"math/rand"
)

// RandomBool returns a random boolean value.
func RandomBool() bool {
	return rand.Intn(2) == 0
}

// RandomInRange returns a pseudorandom integer between min and max (inclusive).
func RandomInRange(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// Abs - find the absolute int (module)
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
