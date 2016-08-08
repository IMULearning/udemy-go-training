package main

import "fmt"

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for n := 0; n < 100000; n++ {
			c <- n
		}
		close(c)
	}()

	go func() {
		for n := range c {
			fmt.Println(n)
		}
		done <- true
	}()

	go func() {
		for n := range c {
			fmt.Println(n)
		}
		done <- true
	}()

	<-done
	<-done
}
