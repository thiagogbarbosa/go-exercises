package main

import "fmt"

/*- Exercício:
  - Crie uma função que retorne um int
  - Crie outra função que retorne um int e uma string
  - Chame as duas funções
  - Demonstre seus resultados */

func sum(a, b int) int {
	return a + b
}

func sumAndSay(a, b int) (string, int) {
	var say string
	say = "Thank you for using our system! Your sum is: "
	return say, a + b

}
func main() {
	fmt.Println(sum(2, 3))
	fmt.Println(sumAndSay(2, 3))

}
