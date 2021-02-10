package main

import (
	"fmt"
)

type person struct {
	name string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hello, my name is", p.name)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{"Alex"}
	saySomething(&p)
	p.speak()
}
