package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type users []user

func (us users) String() {
	for i, u := range us {
		fmt.Printf("#%d - %s %s, %d:\n", i, u.First, u.Last, u.Age)
		sort.Strings(u.Sayings)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}
}

// SortByAge is users alias for sorting
type SortByAge []user

func (a SortByAge) Len() int           { return len(a) }
func (a SortByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// SortByLastname is users alias for sorting
type SortByLastname []user

func (a SortByLastname) Len() int           { return len(a) }
func (a SortByLastname) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByLastname) Less(i, j int) bool { return a[i].Last < a[j].Last }

func main() {

	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := users{u1, u2, u3}

	users.String()

	fmt.Println("----- Sorted by Age -----")

	sort.Sort(SortByAge(users))
	users.String()

	fmt.Println("----- Sorted by Lastname -----")

	sort.Sort(SortByLastname(users))
	users.String()
}
