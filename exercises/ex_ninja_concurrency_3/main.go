package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Practicing Concurrency")
	i := 0
	var wg sync.WaitGroup

	go func() {
		fmt.Println("From Go routine 1")
		myI := i
		runtime.Gosched()
		myI++
		i = myI
		wg.Done()
	}()
	wg.Add(1)

	go func() {
		fmt.Println("From Go routine 2")
		myI := i
		runtime.Gosched()
		myI++
		i = myI
		wg.Done()
	}()
	wg.Add(1)

	fmt.Println(runtime.NumCPU())

	fmt.Println(runtime.NumGoroutine())
	fmt.Println("Waiting...")
	wg.Wait()
	fmt.Println(i)

	FixedRace()
	atomicFix()
	fmt.Println("Exiting Main...")
}
