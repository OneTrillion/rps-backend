package util

import (
	"math/rand"
	"time"
)

// Returns a value between 1 and 3
// where 1 is rock 2 is paper and 3 is scissors
func RandomRpsChoice() int32 {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 3
	rpsChoice := rand.Intn(max-min+1) + min
	return int32(rpsChoice)
}

// Returns a random name from an array
func RandomName() string {
	names := [...]string{
		"John", "Steve", "Amanda", "Benjamin", "Love", "Lucas", "Lara", "Kevin", "Sten",
	}

	rand.Seed(time.Now().UnixNano())
	namePosition := rand.Intn(len(names))
	randName := names[namePosition]
	return randName
}
