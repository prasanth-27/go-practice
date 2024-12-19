package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
)

func main() {
	fmt.Println(`Praticing with:
	 zero value, 
	 short declaration, 
	 multiple initializations, 
	 use var when required, 
	 blank identifier(_)`)

	var a int
	var b bool

	var c, d, e float64

	cat, dog := 0, true

	fmt.Printf("Cat: %d, %T\n", cat, cat)
	fmt.Printf("Dog: %t, %T\n", dog, dog)

	d = math.Pi
	e = rand.Float64()

	fmt.Println("Multiple declared and assigned value to d: ", d)
	fmt.Println("Multiple declared and assigned value to e: ", e)

	fmt.Printf("zero value of int is: %d\n", a)
	fmt.Println("zero value of bool is: ", b)
	fmt.Println("zero value of float64 is: ", c)

	f := 10
	g := 1.2
	h := true
	fmt.Printf("Type and value of short declared f(int) : %T, %v\n", f, f)
	fmt.Printf("Type and value of short declared f(float) : %T, %v\n", g, g)
	n, _ := fmt.Printf("Type and value of short declared f(bool) : %T, %v\n", h, h)
	fmt.Fprintf(os.Stdout, "\nwrote %d charaters in previous write.\n", n)

}
