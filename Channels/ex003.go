package main

import (
	"fmt"
)

//Use a for range loop to demonstrate channel´s value

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

//First step: To create a the receive func that is defined

func receive(canal <-chan int) {
	for v := range canal {
		fmt.Println(v)
	}
}
