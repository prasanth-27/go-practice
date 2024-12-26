package main

import (
	"fmt"
	"time"

	"math/rand"
)

func main() {
	fmt.Println("Practicing variation and infinite for loop")
	i := 10
	for i > 0 {
		fmt.Println(i)
		i--
	}

	num := 1
	for ; num < 10; num++ {
		if num%2 == 0 {
			continue
		}
		fmt.Printf("An odd number %d\n", num)
	}

	for i := 0; i < 5; i++ {
		for j := range 5 {
			fmt.Printf(" iteration %d-->%d\n", i, j)
		}
	}

	sl := 0
	for {
		if sl > 3 {
			break
		}
		fmt.Printf("slept %v seconds\n", sl)
		time.Sleep(time.Second)
		sl++
	}

	xi := []int{42, 43, 44, 45, 46}

	for i, v := range xi {
		fmt.Printf("value:%d at index %d\n", v, i)
	}

	m := map[string]int{
		"first":  42,
		"second": 43,
	}

	for k, v := range m {
		fmt.Printf("value:%d at key %s\n", v, k)
	}

	fmt.Println(m["q"])

	if val, ok := m["1"]; !ok {
		fmt.Println("m['q'] not present in map")
	} else {
		fmt.Println(val)
	}

	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Println("X is 3")
		} else {
			fmt.Println("x is not 3")
		}
	}
	/*
		 	iteration := 0
			for {
				fmt.Println("iteration: ", iteration)
				iteration++
			}
	*/
}
