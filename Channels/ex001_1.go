package main

import (
	"fmt"
)

/*
Faça esse código funcionar utilizando buffer

	func main() {
		c := make(chan int)

		c <- 42

		fmt.Println(<-c)
	}
*/

func main() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)

}
