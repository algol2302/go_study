package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

type personError struct {
	person person
	err    error
}

func (pe personError) Error() string {
	return fmt.Sprintf("a person %v error occured:  %v", pe.person, pe.err)
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		// return []byte{}, errors.New(fmt.Sprintf("Can't create json with error: %v", err))
		return []byte{}, fmt.Errorf("Can't create json for person: %v error: %v", a, err)
	}
	return bs, nil
}
