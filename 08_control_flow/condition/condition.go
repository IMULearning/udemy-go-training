package condition

import "fmt"

func RunCondition() {
	if 10 < 11 {
		fmt.Println("10 is less than 11")
	}

	// A way to keep scope tight
	if ok := isOkay(); ok {
		fmt.Println("Everything is okay")
	}
}

func isOkay() bool {
	return true
}
