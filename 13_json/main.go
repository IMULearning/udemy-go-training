package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	First 	string
	Last 	string	`json:"-"`
	Age 	int 	`json:"active years"`
}

func main() {

	// marshal
	p1 := Person{"James", "Bond", 27}
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))

	// unmarshal
	var p2 Person
	bs2 := []byte(`{"First":"James","Last":"Bond","active years":27}`)
	json.Unmarshal(bs2, &p2)
	fmt.Println(p2.First)
	fmt.Println(p2.Last)
	fmt.Println(p2.Age)

	// encode
	json.NewEncoder(os.Stdout).Encode(&p1)

	// decode
	rdr := strings.NewReader(`{"First":"James","Last":"Bond","active years":27}`)
	var p3 Person
	json.NewDecoder(rdr).Decode(&p3)
	fmt.Println(p3.First)
	fmt.Println(p3.Last)
	fmt.Println(p3.Age)
}
