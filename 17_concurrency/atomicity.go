package main

import (
	"sync"
	"time"
	"math/rand"
	"fmt"
	"sync/atomic"
)

var wg sync.WaitGroup
var counter int64

func main() {
	wg.Add(2)
	go incrementor3("Foo:")
	go incrementor3("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor3(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}