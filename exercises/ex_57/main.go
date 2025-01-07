package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println("Practicing function concepts!")

	fmt.Println(sum([]int{1, 2, 3}))

	fmt.Printf("Func foo returned: %d\n", foo())

	i, s := bar()
	fmt.Printf("bar() return : %d, %s\n", i, s)

	si := make([]int, 10)
	for i := range 10 {
		si[i] = rand.Intn(1000)
	}
	fmt.Println(si)
	fmt.Printf("sum from sum([]int): %d \nSum from variadicSum(xi... int): %d\n", sum(si), variadicSum(si...))

	deferTest()

	james := person{
		name: "James Bond",
		age:  35,
	}

	james.speak()
}

// named return
// return var is already declared in the signature
func sum(xi []int) (total int) {
	for _, v := range xi {
		total += v
	}
	return
}

// no parameters and one return
func foo() int {
	i := rand.Int()
	return i
}

// no parameters multiple returns
func bar() (int, string) {
	i := rand.Int()
	s := strconv.FormatFloat(rand.Float64(), 'e', 5, 64)
	return i, s
}

// Accepts variable count of arguments into a slice int xi
func variadicSum(xi ...int) int {
	t := 0
	for _, v := range xi {
		t += v
	}
	return t
}

// defer func : executed after returning
// defer executed LIFO
func deferTest() {
	defer fmt.Println("deferTest Exited...")
	defer fmt.Println(3)
	fmt.Println("Invoked deferTest")
}

// Methods /functions with receiver
func (p person) speak() {
	fmt.Printf("Hello! I am %s and I am %d old.\n", p.name, p.age)
}
