package main

import "fmt"

func main() {
	fmt.Println("Practicing Tests in go!")

	fmt.Println("====================")
	fmt.Println(Add(3, 2))
	fmt.Println(doMath(3, 2, Add))

	fmt.Println("====================")
	fmt.Println(substract(3, 2))
	fmt.Println(doMath(3, 2, substract))
}

func Add(x, y int) int {
	return x + y
}

func substract(x, y int) int {
	return x - y
}

func doMath(x, y int, op func(int, int) int) int {
	return op(x, y)
}
