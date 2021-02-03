package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	// people := []person{}
	var people []person
	fmt.Println(people)

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(people)

	for i, v := range people {
		fmt.Println("Person(", i, "):", v.First, v.Last, "age:", v.Age)
	}
}
