package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Practicing Concurrency")

	var wg sync.WaitGroup

	go func() {
		fmt.Println("From Go routine 1")
		wg.Done()
	}()
	wg.Add(1)

	go func() {
		fmt.Println("From Go routine 2")
		wg.Done()
	}()
	wg.Add(1)

	fmt.Println(runtime.NumCPU())

	fmt.Println(runtime.NumGoroutine())
	fmt.Println("Waiting...")
	wg.Wait()
	fmt.Println("Exiting Main...")
}
