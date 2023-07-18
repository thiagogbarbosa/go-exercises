package main

//Exercício proposto: remova o underscore e verifique e lide com o erro de maneira apropriada.

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

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Println("There is an err occuring", err)
	}
	fmt.Println(string(bs))

}
