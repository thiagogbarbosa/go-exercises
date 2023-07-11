package main

import (
	"fmt"
)

/*
Faça esse código funcionar utilizando uma func anônima

	func main() {
		c := make(chan int)

		c <- 42

		fmt.Println(<-c)
	}
*/
func main() {
	c := make(chan int)

	go func() {
		c <- 42

	}()
	fmt.Println(<-c)

}
