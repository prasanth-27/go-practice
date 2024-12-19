package main

import "fmt"

type memsize int

const (
	_ memsize = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
	EB
)

func main() {
	fmt.Println("Printing Sized of memory using iota!!")
	fmt.Printf("Size KB: %+v \t %b\n", KB, KB)
	fmt.Printf("Size MB: %#v \t %b\n", MB, MB)
	fmt.Printf("Size GB: %#v \t %b\n", GB, GB)
	fmt.Printf("Size TB: %#v \t %b\n", TB, TB)
	fmt.Printf("Size PB: %#v \t %b\n", PB, PB)
}
