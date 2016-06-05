package main

import "fmt"

func zero(x *int) {
	*x = 0
}

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 42
	fmt.Println(*b)
	fmt.Println(a)

	x := 5
	zero(&x)
	fmt.Println(x)
}
