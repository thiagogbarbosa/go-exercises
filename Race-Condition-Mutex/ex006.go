package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*Corringo a Condição de corrida com Mutex*/

var wg sync.WaitGroup
var contador int
var mu sync.Mutex

const quantidadeDeGoroutines = 300

func goroutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()
	}
}

func main() {
	goroutines(quantidadeDeGoroutines)
	wg.Wait()
	fmt.Println("Total de goroutines:\t", quantidadeDeGoroutines, "\nTotal do contador:\t", contador)

}
