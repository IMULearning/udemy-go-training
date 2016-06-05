package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

// not gonna use it? use _
func main() {
	res, _ := http.Get("http://www.google.ca")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
