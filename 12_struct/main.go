package main

import "fmt"

type Person struct {
	First string
	Last string
	Age int
}

type Agent struct {
	Person
	Weapon string
}

func (p *Person) greeting() string {
	return "Hello World"
}

func (p Person) fullName() string {
	return p.First + p.Last
}

func (p Agent) fullName() string {
	return "Agent: " + p.Person.fullName()
}

type foo int

func main() {

	// person
	p1 := &Person{ "David", "Qiu", 27}
	p2 := &Agent {
		Person {
			First: "James",
			Last: "Bond",
			Age: 45,
		},
		"glock",
	}

	fmt.Printf("%T\n", p1)

	fmt.Println(p1.First, p1.Last, p1.Age)
	fmt.Println(p2.First, p2.Last, p2.Age, p2.Weapon)

	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
	fmt.Println(p2.Person.fullName())

	p3 := &Person{"David", "Qiu", 27}
	fmt.Println(p3.greeting())
}
