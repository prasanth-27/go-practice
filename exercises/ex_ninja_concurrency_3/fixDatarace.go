package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func FixedRace() {
	fmt.Println("In fixed race")
	i := 0
	var wg sync.WaitGroup
	var m sync.Mutex

	go func() {
		fmt.Println("From Go routine 1")
		m.Lock()
		myI := i
		runtime.Gosched()
		myI++
		i = myI
		m.Unlock()
		wg.Done()
	}()
	wg.Add(1)

	go func() {
		fmt.Println("From Go routine 2")
		m.Lock()
		myI := i
		runtime.Gosched()
		myI++
		i = myI
		m.Unlock()
		wg.Done()
	}()
	wg.Add(1)

	fmt.Println(runtime.NumCPU())

	fmt.Println(runtime.NumGoroutine())
	fmt.Println("Waiting...")
	wg.Wait()
	fmt.Println(i)
	fmt.Println("Exiting fixed...")
}

func atomicFix() {
	fmt.Println("InAtomicFix")
	//var i int64 = 0
	var i atomic.Int64
	var mwg sync.WaitGroup
	mwg.Add(100)
	for itr := 0; itr < 100; itr++ {

		go func() {
			//i = atomic.AddInt64(&i, int64(1))
			i.Add(1)
			mwg.Done()
		}()
	}
	mwg.Wait()
	fmt.Println("End value:", i.Load())
}
