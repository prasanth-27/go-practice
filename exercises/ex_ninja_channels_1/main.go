package main

import "fmt"

func main() {
	fmt.Println("Practicing channels!")

	c := make(chan int)

	// c <- 42 // This doesn't work as it will block
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)

	//Buffered channel can hold 1 item before blocking
	c1 := make(chan int, 1)
	c1 <- 43
	fmt.Println(<-c1)

	g := gen()

	for v := range g {
		fmt.Println(v)
	}

	fmt.Println("===================")

	r := make(chan int)
	g = gen2(r)
	receive(g, r)

	fmt.Println("===================")

	c = make(chan int)

	go func() {
		c <- 42

	}()

	v, ok := <-c
	fmt.Println(v, ok)
	close(c)
	fmt.Println("==")
	v, ok = <-c
	fmt.Println(v, ok)

}

func gen() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func gen2(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		q <- 1
	}()
	return c
}

func receive(g, r <-chan int) {
	for {
		select {
		case v := <-g:
			fmt.Println(v)
		case v := <-r:
			fmt.Println("Exiting select", v)
			return
		}
	}
}
