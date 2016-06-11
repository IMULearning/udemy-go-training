package main

import "fmt"

func hello() {
	fmt.Print("hello ")
}

func world() {
	fmt.Println("world")
}

func main() {
	defer world() 	// defer to before main() exits
	hello()
}
