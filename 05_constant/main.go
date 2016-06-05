package main

import "fmt"

const p string = "Hello!"

const (
	PI = 3.14
	E  = 2.71
)

func main() {
	const q = 42

	fmt.Println("p: ", p)
	fmt.Println("q: ", q)

	fmt.Println(PI)
	fmt.Println(E)

}
