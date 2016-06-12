package main

import "fmt"

type array []int

func (a array) len() int {
	return len(a)
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(array(a).len())
}
