package main

import "fmt"

func main() {
	fmt.Println("Practicing Arrays and slices #2 !")

	arrI := [5]int{}
	arrI[0] = 1
	arrI[1] = 2
	arrI[2] = 3
	arrI[3] = 4
	arrI[4] = 5

	fmt.Println(arrI)
	fmt.Printf("Type of arrI: %T\n", arrI)
	fmt.Println(len(arrI))
	fmt.Println(cap(arrI))

	for i, v := range arrI {
		fmt.Printf("Values:%d, at index:%d\n", v, i)
	}

	xi := []int{}
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	for i := 0; i < 10; i++ {
		fmt.Println("Iter-----")
		xi = append(xi, i)
		fmt.Printf("%p\n", xi)
		fmt.Println(len(xi))
		fmt.Println(cap(xi))
	}

	for _, v := range xi {
		fmt.Println(v)
	}

	fmt.Println(xi[:])
	fmt.Println(xi[:1])
	fmt.Println(xi[:5])
	fmt.Println(xi[1:])
	fmt.Println(xi[3:7]) // Index 3 to index 7(exclusive)

	fmt.Println("=======================")

	xi = append(xi, 52)
	fmt.Println(xi)

	xi = append(xi, 53, 54, 55)
	fmt.Println(xi)

	y := []int{56, 57, 58, 59, 60}
	fmt.Println(xi, y)
	xi = append(xi, y...)
	fmt.Println(xi)

	fmt.Println("==================")

	x := append(xi[7:10], xi[12:15]...)
	fmt.Println(x)

	//No builtin to delete element from slice
	//Should use slicing
	x = append(x[:2], x[3:]...)
	fmt.Println(x)

	fmt.Println("================")

	ci := make([]int, 10, 20)
	fmt.Println(ci, len(ci), cap(ci))

	for i := range 10 {
		fmt.Printf("Iteration:%d", i+1)
		fmt.Printf("slice:%#v \t address:%p \t length:%d \t capacity:%d\n", ci, &ci, len(ci), cap(ci))
		ci = append(ci, i)
		ci[i] = i
		fmt.Printf("slice:%#v \t length:%d \t capacity:%d\n", ci, len(ci), cap(ci))
	}

	xs := make([][]string, 2)
	fmt.Printf("val:%#v, \t type:%T\n", xs, xs)
	xs = append(xs, []string{"First", "slice"})
	fmt.Printf("val:%#v, \t type:%T\n", xs, xs)

	s := []string{"Second", "slice"}
	xs = append(xs, s)
	fmt.Printf("val:%#v, \t type:%T\n", xs, xs)

}
