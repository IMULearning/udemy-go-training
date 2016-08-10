package main

import (
	"log"
	"github.com/go-errors/errors"
)

func main() {
	_, err := Sqrt(-10)
	if (err != nil) {
		log.Fatalln(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("negative!")
	}
	return 42, nil
}
