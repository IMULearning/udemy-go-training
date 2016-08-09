package main

import "fmt"

func main() {
	c := supplier(5)
	c0 := consumer(c)
	fmt.Println(<-c0)
}

func supplier(num int) chan int {
	c := make(chan int)
	go func() {
		for i := 1; i <= num; i++ {
			c <- i
		}
		close(c)
	}()
	return c;
}

func consumer(c chan int) chan int {
	c0 := make(chan int)
	go func() {
		var fac = 1
		for n := range c {
			fac *= n
		}
		c0 <- fac
		close(c0)
	}()
	return c0
}
