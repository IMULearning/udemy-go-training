package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []string{"v", "as", "q"}
	sort.Sort(sort.StringSlice(a))
	fmt.Println(a)
}
