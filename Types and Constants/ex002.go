package main

import "fmt"

/*Crie constantes tipadas e não tipadas*/

const eNumber float64 = 2.718281828459045235
const notAnumber = "This is not a number"

func main() {
	fmt.Println(eNumber, notAnumber)
	fmt.Printf("Type of %v is %T\nType of %v is %T", eNumber, eNumber, notAnumber, notAnumber)

}
