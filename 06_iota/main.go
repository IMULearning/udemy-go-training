package main

import "fmt"

const (
	_  = iota		// throw away 0
	KB = 1 << (iota * 10) 	// 1 << (1 * 10)
	MB = 1 << (iota * 10)	// 1 << (2 * 10)
)

func main() {
	fmt.Println(KB)
	fmt.Println(MB)
}
