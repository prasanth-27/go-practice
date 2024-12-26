package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	fmt.Println("Practicing switch statement")

	num := rand.Intn(250)

	fmt.Fprintf(os.Stdout, "Random number generated is: %d\n", num)

	switch {
	case num <= 100:
		fmt.Println("Number less than equal to 100")

	case num > 100 && num <= 200:
		fmt.Println("Number greater than 100 and less than equal to 200")

	case num > 200 && num < 250:
		fmt.Println("Number greater than 200 and less than equal to 250")

	}

}
