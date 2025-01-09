package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(`Practicing:
		Anonymus funcs
		func expressions
		func return func
		callback
		closure
		wrapper
	`)

	sum := func(x, y int) int {
		fmt.Println("Inside anon func")
		return x + y
	}(3, 2)

	fmt.Printf("%#v\n", sum)

	fmt.Println("======================")

	logger := func() {
		fmt.Println("Logger invoked")
	}

	logger()

	fmt.Println("======================")

	summer := func() func(int, int) int {
		return add
	}

	fmt.Printf("type of summer:%#v, summer():%#v, return :%v", summer, summer(), summer()(3, 2))

	fmt.Println("======================")

	fmt.Println(printSquare(square, 2))

	fmt.Println("======================")
	counter := incrementer()

	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	fmt.Println("============")

	timeFunc(func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
	})
}

func add(x, y int) int {
	return x + y
}

func square(x int) int {
	return x * x
}

func printSquare(s func(int) int, num int) string {
	return fmt.Sprintf("Square of %d is %d", num, s(num))
}

func incrementer() func() int {
	var i int
	return func() int {
		i++
		return i
	}
}

func timeFunc(f func()) {
	start := time.Now()

	f()

	end := time.Now()

	fmt.Printf("func took %v", end.Sub(start))
}
