package main

import (
	"fmt"

	"math/rand"
)

func main() {
	fmt.Println("Practicing for loop")

	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}

	i := 1
	for ; i <= 42; i++ {
		num := rand.Intn(5)

		fmt.Printf("Iteration %d generated %d", i, num)
		switch num {
		case 0:
			fmt.Println("Number is 0")
		case 1:
			fmt.Println("Number is 1")
		case 2:
			fmt.Println("Number is 2")
		case 3:
			fmt.Println("Number is 3")
		case 4:
			fmt.Println("Number is 4")
		default:
			fmt.Println("Number greater than 4")
		}
	}

}
