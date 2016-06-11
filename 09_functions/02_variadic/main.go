package main

import "fmt"

func main() {
	fmt.Println(average(3, 5, 7))

	sf := []float64{3, 5, 6, 7, 8}
	fmt.Println(average(sf...))
}

func average(sf ...float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
