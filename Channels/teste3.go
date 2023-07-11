package main

import (
	"fmt"
	"sync"
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
	var wg sync.WaitGroup
	totalgoroutines := (1)

	go func() {
		wg.Add(totalgoroutines)
		c <- 42
		fmt.Println(<-c)
		wg.Done()
	}()

	wg.Wait()
	sync.WaitGroup()
}
