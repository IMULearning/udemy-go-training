package rune

import "fmt"

func LoopRune() {
	for i := 50; i <= 140; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
}
