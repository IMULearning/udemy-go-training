package main

import (
	"sync"
	"time"
	"math/rand"
	"fmt"
)

var wg sync.WaitGroup
var mutex sync.Mutex
var counter int

func main() {
	wg.Add(2)
	go incrementor2("Foo:")
	go incrementor2("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor2(s string) {
	for i := 0; i < 20; i++ {
		mutex.Lock()
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter++
		fmt.Println(s, i, "Counter:", counter)
		mutex.Unlock()
	}
	wg.Done()
}