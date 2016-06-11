package main

import "fmt"

func main() {
	visit([]int{1, 2, 3, 4, 5}, func (n int) {
		fmt.Println(n)
	})

	fmt.Println(filter([]int{1, 2, 3, 4, 5}, func (n int) bool {
		return n > 3
	}));
}

func visit(numbers []int, callback func (int)) {
	for _, n := range numbers {
		callback(n)
	}
}

func filter(numbers []int, criteria func(int) bool) []int {
	result := []int{}
	for _, n := range numbers {
		if criteria(n) {
			result = append(result, n)
		}
	}
	return result
}
