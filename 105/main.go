package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("I'm", p.first, p.last, "and I'm", p.age, "years old")
}

func main() {
	p := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p.speak()
}
