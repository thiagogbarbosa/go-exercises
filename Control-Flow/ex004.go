package main

import "fmt"

/* Demonstre o resto da divisão por 4 de todos os números entre 10 e 100*/

func main() {
	for i := 10; i <= 100; i++ {
		restoDivisão := i % 4
		fmt.Printf("O resto da divisão de %v por 4 é %v\n", i, restoDivisão)

	}

}
