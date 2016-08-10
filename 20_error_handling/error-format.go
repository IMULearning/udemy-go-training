package main

import "fmt"

func main() {
	err := fmt.Errorf("%s is wrong", "something")
	fmt.Println(err)
}
