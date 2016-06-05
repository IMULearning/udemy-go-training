package main

import "fmt"

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())

	incrementWrapper := wrapper()
	fmt.Println(incrementWrapper())
	fmt.Println(incrementWrapper())
}
