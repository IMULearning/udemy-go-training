package main

import "fmt"

func main() {
	greeter := makeGreeter()
	fmt.Println(greeter("David"))
}

func makeGreeter() func(name string) string {
	return func(name string) string {
		return "Hello " + name
	}
}
