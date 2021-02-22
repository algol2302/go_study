package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()

	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")
	if err != nil {
		//		fmt.Println("err happened", err)
		// log.Println("err happened 2", err)
		// log.Fatalln(err) // defered foo doesn't run
		log.Panicln(err) // defered foo runs
		// panic(err)
	}
	defer f2.Close()

	fmt.Println("check the log.txt file in the directory")
}

func foo() {
	fmt.Println("Deferred text")
}

// Println calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.
