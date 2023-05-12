package Utils

import (
	"math/rand"
	"time"
)

func RandomNumber() (number int){
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 0 and 100
	randomNumber := rand.Intn(10000)
	return randomNumber
}