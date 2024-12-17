package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func anotherwayadd(x, y int) int {
	return x + y
}

func swap(x, y int) (int, int) {
	return y, x
}
func anotherswap(x, y int) (a, b int) {
	a, b = y, x
	return
	// this is a naked return as the vars are already in return part
}

func main() {
	fmt.Println(add(1, 2))

	x, y := 2, 3
	a, b := swap(x, y)
	fmt.Printf("Swapping (%d,%d): (%d,%d)\n", x, y, a, b)

	a, b = anotherswap(x, y)
	fmt.Printf("another Swapping (%d,%d): (%d,%d)\n", x, y, a, b)

	//variaqbles

	var q int = 10
	var w = 10
	fmt.Println("both are valid", q, w)
	fmt.Println(":= construct is not available outside the function socpe")
	builtInDtypes := `
  bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128`

//factored var declaration

var(
	a:=10
	b:=20
	c:="hello"
	f:=false
	d:=2+3i //complex datatype
)

flt:=float64(b) //converions

const pi = math.Pi


}
