package main

import "fmt"

/*Crie um programa que demonstre o funcionamento da declaração if.*/

func main() {

	x := 10

	if x%2 == 0 {
		fmt.Printf("O número %v é par", x)
	} else {
		fmt.Printf("O número %v é impar!", x)

	}

}
