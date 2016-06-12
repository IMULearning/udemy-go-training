package main

import "fmt"

func main() {
	mySlice := []string{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
		"g",
	}
	fmt.Println(mySlice[2:])
	fmt.Println(mySlice[:4])
	fmt.Println(mySlice[2:4])
	fmt.Println(mySlice[:])

	a := make([]string, 0, 50)
	fmt.Println(a)

	b := make([]int, 0, 5)
	fmt.Println(b)

	for i := 0; i < 80; i++ {
		b = append(b, i)
		fmt.Println("Len:", len(b), "Capacity:", cap(b), "Value:", b[i])
	}
}
