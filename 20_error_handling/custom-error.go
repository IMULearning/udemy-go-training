package main

import "fmt"

type CustomError struct {
	code	string
	err	error
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Code is %s", e.code)
}

func main() {
	_, err := somethingWrong()
	if err != nil {
		fmt.Println(err)
	}
}

func somethingWrong() (int, error) {
	return 0, &CustomError{code: "12345"}
}
