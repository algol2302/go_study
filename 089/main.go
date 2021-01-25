package main

import "fmt"

func main() {
	anonymousStruct := struct {
		field1 []string
		field2 map[string]string
		field3 int
	}{
		field1: []string{"A", "B"},
		field2: map[string]string{"a": "aaaa", "b": "bbbb"},
		field3: 3,
	}
	fmt.Println(anonymousStruct)
}
