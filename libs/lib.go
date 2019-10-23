package libs

import "math/rand"

func GetRandInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}
