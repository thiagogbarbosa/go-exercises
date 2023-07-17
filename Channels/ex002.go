package main

import (
	"fmt"
)

//Make this code works

func main() {
	//Original code
	//cs := make(chan<- int)
	cs := make(chan int) //Just need to remove the receiver signal

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
