package main

import "fmt"

func main() {
	x := map[string][]string{
		"bond_james":      {`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": {`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	x["test"] = []string{"test1", "test2", "test3"}

	delete(x, "no_dr")

	for key, values := range x {
		fmt.Println("key: ", key)
		for index, value := range values {
			fmt.Printf("\t index: %v \t value: %v\n", index, value)
		}
	}
}
