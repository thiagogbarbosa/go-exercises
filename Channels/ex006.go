package main

import "fmt"

// Escreva um programa que coloque 100 números em um canal, retire os números do canal, e demonstre-os.

func main() {
	c := make(chan int)
	gen(c)
	dem(c)
}

func gen(canal chan int) chan<- int {
	go func() {
		for i := 0; i < 100; i++ {
			canal <- i
		}
		close(canal)

	}()
	return canal
}

func dem(canal chan int) <-chan int {

	for v := range canal {
		fmt.Println(v)
	}
	return canal
}
