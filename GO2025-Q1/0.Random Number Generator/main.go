package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello")
	// Create a new source of randomness
	source := rand.NewSource(time.Now().UnixNano())
	// Create a new random number generator
	r := rand.New(source)

	// Generate Integer random integer between 0 (inclusive) and 100 (exclusive)
	fmt.Println(r.Intn(100))

	// Generate a random integer between 10 and 20 (inclusive)
	min := 10
	max := 20
	getRandVal := r.Intn(max-min+1) + min
	fmt.Println(getRandVal)

}
