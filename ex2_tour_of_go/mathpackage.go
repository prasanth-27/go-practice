package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("A rondom number between 1 and 10 exclusive of 10 is :", rand.Intn(10))
	fmt.Println("Squareroot of 7 is: ", math.Sqrt(7))

	fmt.Println("value of Pi(but not pi) is: ", math.Pi)
}
