package main

import (
	"os"
	"fmt"
	"log"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("error happened:", err)
		log.Println("error happened:", err)
		log.Fatalln("error happened:", err)
		panic(err)
	}
}
