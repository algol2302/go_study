package main

import (
	"fmt"
)

func main() {
	year := 1981
	for {

		if year > 2021 {
			break
		}

		fmt.Println(year)
		year++
	}
}
