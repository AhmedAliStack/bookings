package helpers

import "math/rand"

func CreateRandom(a int) int {
	return rand.Intn(a)
}