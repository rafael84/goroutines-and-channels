package random

import "math/rand"

func IntBetween(min int, max int) int {
	return min + rand.Intn(max-min)
}
