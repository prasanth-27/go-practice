package ex19

import (
	"fmt"
	"math/rand"
)

func Ex19() {

	num := rand.Intn(250)

	if num > 0 && num <= 100 {
		fmt.Println("Number is between 0 and 100")
	} else if num > 100 && num <= 200 {
		fmt.Println("Number is between 101 and 200")
	} else if num > 200 && num <= 250 {
		fmt.Println("Number is between 201 and 250")
	}
}
