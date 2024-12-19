package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Practicing format verbs binary,decimal,hex")

	const (
		a = 747
		b = 911
		d = 9021
		e = math.Pi
	)

	fmt.Printf("%d\t %b\t %o\t %x\n", a, a, a, a)
	fmt.Printf("%d\t %b\t %o\t %X\n", b, b, b, b)
	fmt.Printf("%d\t %b\t %o\t %#X\n", d, d, d, d)
	fmt.Printf("%f\t %3f\t %.3f\n", e, e, e)
}
