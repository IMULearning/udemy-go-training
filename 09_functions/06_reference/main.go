package main

import "fmt"

func main() {
	a := 24
	changeInt(&a)
	fmt.Println(a)

	// make already makes reference type, so no need to pass address, it is already an pointer
	s := make([]string, 1, 25)
	fmt.Println(s)
	changeSlice(s)
	fmt.Println(s)
}

func changeInt(a *int) {
	fmt.Println(*a)
	*a = 44
	fmt.Println(*a)
}

func changeSlice(s []string) {
	s[0] = "David"
	fmt.Println(s)
}
