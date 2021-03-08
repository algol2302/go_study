package main

import (
	"fmt"

	"github.com/algol2302/go_study/190/quote"
	"github.com/algol2302/go_study/190/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	fmt.Println(len(quote.SunAlso))
	fmt.Println(len(word.UseCount(quote.SunAlso)))

	// for k, v := range word.UseCount(quote.SunAlso) {
	// 	fmt.Println(v, k)
	// }
}
