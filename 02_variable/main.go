package main

import "fmt"

func main() {

	a := 10
	b := "golang"
	c := 4.17
	d := false

	fmt.Printf("%v (%T) \n", a, a)
	fmt.Printf("%v (%T) \n", b, b)
	fmt.Printf("%v (%T) \n", c, c)
	fmt.Printf("%v (%T) \n", d, d)

	var e int
	var f string
	var g float64
	var h bool

	fmt.Printf("%v (%T) \n", e, e)
	fmt.Printf("%v (%T) \n", f, f)
	fmt.Printf("%v (%T) \n", g, g)
	fmt.Printf("%v (%T) \n", h, h)
}
