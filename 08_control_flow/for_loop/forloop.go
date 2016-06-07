package for_loop

import "fmt"

func RunLoopType1() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func RunLoopType2() {
	x := 0
	for x < 5 {
		fmt.Println(x)
		x++
	}
}

func RunLoopType3() {
	x := 0
	for {
		x++
		if x%2 == 0 {
			continue
		}

		fmt.Println(x)
		if x >= 10 {
			break
		}
	}
}
