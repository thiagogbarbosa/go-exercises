package main

import "fmt"

/* Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
- Passe um valor do tipo slice of int como argumento para a função;
- Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
- Passe um valor do tipo slice of int como argumento para a função.*/

//Variadic parameter
func sumSlice(x ...int) int {
	total := 0
	for _, v := range x {
		total += v

	}
	return total
}

// Slice of int as an parameter
func sumSlice2(y []int) int {
	total2 := 0
	for _, num := range y {
		total2 += num
	}
	return total2
}

var oddSlice []int
var evenSlice []int

func main() {

	oddSlice = []int{1, 3, 5, 7, 9}
	evenSlice = []int{2, 4, 6, 8, 10}
	fmt.Println(sumSlice(evenSlice...)) //When using slice of int as an argument for an int parameter its needed the variadic term "..."
	fmt.Println(sumSlice2(oddSlice))
}
