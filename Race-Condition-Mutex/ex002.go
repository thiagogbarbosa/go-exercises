package main

import (
	"fmt"
	"runtime"
	"sync"
)

//Resolvendo o problema do Código em que h+a condição de corrida
//Utilizando o Mutex

func main() {

	fmt.Println("CPUs:", runtime.NumCPU()) //somente uma curiosidade mesmo
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	contador := 0
	totaldegoroutines := 15

	var wg sync.WaitGroup
	wg.Add(totaldegoroutines)

	var mu sync.Mutex

	for i := 0; i < totaldegoroutines; i++ {
		go func() {
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}() //para executar a func anonima
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Valor final:", contador)
}
