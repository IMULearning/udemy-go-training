package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	delete(m, "k2")
	fmt.Println("map:", m)

	v, ok := m["k1"]
	fmt.Println("ok:", ok, v)

	n := map[string]int {
		"a": 1,
		"b": 2,
	}
	fmt.Println(n)

	if value, exists := m["k2"]; exists {
		fmt.Println(value)
	} else {
		fmt.Println("k2 does not exist")
	}

	for key, val := range m {
		fmt.Println(key, val)
	}
}