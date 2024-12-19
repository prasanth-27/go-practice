//Practicing with iota

package main

import "fmt"

const (
	a = iota
	b
	c
	d
)
const (
	v1 = iota + 10
	v2
	v3
)

func main() {
	fmt.Println("Practicing iota")

	fmt.Println("Value of a is ", a)
	fmt.Println("Value of b is ", b)
	fmt.Println("Value of c is ", c)
	fmt.Println("Value of d is ", d)

	fmt.Println("Value of v1 is ", v1)
	fmt.Println("Value of v2 is ", v2)
	fmt.Println("Value of v3 is ", v3)

	fmt.Printf("%d, \t %b\n", 1, 1)
	fmt.Printf("%d, \t %b\n", 1<<1, 1<<1)
	fmt.Printf("%d, \t %b\n", 1<<2, 1<<2)

	fmt.Printf("%d \t %b\n", 5, 5)
	fmt.Printf("%d \t %b\n", 5>>1, 5>>1)

}
