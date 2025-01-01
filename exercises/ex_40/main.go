package main

import "fmt"

func main() {
	fmt.Println("Practicing arrays.")

	as := [...]string{"Hello", "World", "THis ", "is", "an", "array", "of", "strings!"}

	fmt.Println(as)
	fmt.Println(len(as))

	fmt.Printf("%#v\n", as)
	fmt.Printf("%T\n", as)

	for _, v := range as {
		fmt.Println(v)
	}

}
