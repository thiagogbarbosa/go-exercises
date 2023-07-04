package main

import (
	"fmt"
	"runtime"
	"sync"
)

//Replicando um programa que irá gerar condição de corrida

func main() {

	fmt.Println("CPUs:", runtime.NumCPU()) //somente uma curiosidade mesmo
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	contador := 0
	totaldegoroutines := 1000

	var wg sync.WaitGroup
	wg.Add(totaldegoroutines)

	for i := 0; i < totaldegoroutines; i++ {
		go func() {
			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
		}() //para executar a func anonima
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Valor final:", contador)
}

/*O compartilhamento de variável entre as funções gerou condição de corrida */
