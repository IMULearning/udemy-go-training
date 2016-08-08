package main

import (
	"fmt"
	"time"
)

// a.k.a buffer
// c <- i, code stops until <-c
// once passed to the channel, the code blocks until the value is taken off the channel

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()

	time.Sleep(time.Second)
}
