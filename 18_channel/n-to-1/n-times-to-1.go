package main

import "fmt"

func main() {
	c := make(chan int)
	done := make(chan bool)

	for n := 0; n < 10; n++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
			done <- true
		}()
	}

	go func() {
		for n := 0; n < 10; n++ {
			<-done
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
