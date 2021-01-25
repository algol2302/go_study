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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for key, value := range m {
		fmt.Println(key)
		fmt.Println(value.first, value.last)
		for index, flavor := range value.flavors {
			fmt.Println(index, flavor)
		}
	}

}
