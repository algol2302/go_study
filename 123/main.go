package main

import (
	"fmt"
	"sort"
)

type person struct {
	first string
	age   int
}

func (p person) String() string {
	return fmt.Sprintf("%s: %d", p.first, p.age)
}

// ByAge is helper struct
type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }

// ByName is helper struct
type ByName []person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].first < a[j].first }

func main() {
	p1 := person{
		first: "James",
		age:   32,
	}
	p2 := person{
		first: "Miss",
		age:   27,
	}
	p3 := person{
		first: "Q",
		age:   64,
	}
	p4 := person{
		first: "M",
		age:   56,
	}

	people := []person{p1, p2, p3, p4}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)
}
