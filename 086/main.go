package main

import "fmt"

type person struct {
	first   string
	last    string
	flavors []string
}

func main() {
	p1 := person{
		first:   "James",
		last:    "Bond",
		flavors: []string{"vanilla", "chocolate"},
	}

	p2 := person{
		first:   "Miss",
		last:    "Moneypenny",
		flavors: []string{"banana", "strawberry"},
	}

	fmt.Println(p1.first, p1.last)
	for index, flavor := range p1.flavors {
		fmt.Println(index, flavor)
	}

	fmt.Println(p2.first, p2.last)
	for index, flavor := range p2.flavors {
		fmt.Println(index, flavor)
	}

}
