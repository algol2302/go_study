package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {
	p.age++
	// (*p).age++
}

func main() {
	p := person{
		first: "James",
		last:  "Bond",
		age:   31,
	}
	fmt.Println("before:", p)
	changeMe(&p)
	fmt.Println("after:", p)

}
