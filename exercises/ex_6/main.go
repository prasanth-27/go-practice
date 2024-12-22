package main

import "fmt"

var s string = "This is a global variable of this package."

const i int = 10

func main() {
	fmt.Println("Practicing scope of variables.")

	sL := "This is a local varibale with scope in the main function."

	fmt.Println(s, i)
	fmt.Println(sL)

}
