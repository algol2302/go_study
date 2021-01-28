package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I'm a secret agent", s.first, s.last)
}

func (p person) speak() {
	fmt.Println("I'm a person", p.first, p.last)
}

type human interface {
	speak()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("Person was passed into bar", h.(person).first)
	case secretAgent:
		fmt.Println("Secret Agent was passed into bar", h.(secretAgent).first)
	}

}

func main() {

	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: false,
	}

	p1 := person{
		"Dr.",
		"Yes",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)

	bar(sa1)
	bar(sa2)
	bar(p1)

}
