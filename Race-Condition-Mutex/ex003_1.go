package main

import (
	"fmt"
	"sync"
)

//Mesmo exercício ex003- Porém, mais simples

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go func() {
		fmt.Println("Eu sou a goroutine número: 1")
		wg.Done()
	}()
	go func() {
		fmt.Println("Eu sou a goroutine número: 2")
		wg.Done()
	}()
	wg.Wait()
}
