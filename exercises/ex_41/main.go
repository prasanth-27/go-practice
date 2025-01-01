package main

import "fmt"

func main() {
	fmt.Println("Practicing slices.")

	as := []string{"Hello", "World", "THis ", "is", "a", "slice", "of", "strings!"}

	fmt.Println(as)
	fmt.Println(len(as))

	fmt.Printf("%#v\n", as)
	fmt.Printf("%T\n", as)

	as = append(as, "appended!")

	fmt.Println(len(as))
	fmt.Println(cap(as))

	for _, v := range as {
		fmt.Println(v)
	}

}
