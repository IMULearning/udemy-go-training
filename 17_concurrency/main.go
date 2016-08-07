package main

import (
	"fmt"
	"sync"
	"time"
	"runtime"
)

var wg sync.WaitGroup

// init is a function to setup
// You no longer need to include this after GO 1.5
// Go >1.5 automatically default to use more than 1 cores
// Here we say use all calls
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(time.Duration(3 * time.Microsecond))
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(time.Duration(20 * time.Microsecond))
	}
	wg.Done()
}
