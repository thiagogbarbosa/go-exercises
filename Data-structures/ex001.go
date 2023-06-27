package main

import "fmt"

/*Usando uma literal composta:
    - Crie um array que suporte 5 valores to tipo int
    - Atribua valores aos seus índices
Utilize range e demonstre os valores do array.
Utilizando format printing, demonstre o tipo do array.*/

func main() {
	num := [5]int{1, 2, 3, 4, 5}

	for _, v := range num {
		fmt.Println(v)
	}
	fmt.Printf("Array type: %T", num)

}
