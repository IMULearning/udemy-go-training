package main

import "fmt"

func main() {
	greet("David", "Qiu")
	fmt.Println(greet2("David", "Qiu"))

	s1, s2 := greet3("David", "Qiu")
	fmt.Println(s1, s2)
}

func greet(fname string, lname string) {
	fmt.Println("Hello", fname, lname)
}

func greet2(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}

/*
 Multiple Return
 */
func greet3(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}
