package main

import "fmt"

/*- Crie um struct "erroEspecial" que implemente a interface builtin.error.
- Crie uma função que tenha um valor do tipo error como parâmetro.
- Crie um valor do tipo "erroEspecial" e passe-o para a função da instrução anterior.*/

type erroEspecial struct {
	qualquercoisa string
}

func (e erroEspecial) Error() string {
	return "There is an error"
}

func erroComoParametro(e error) {
	fmt.Println("Foi passado como argumento: ", e)
}

func main() {
	testandoErro := erroEspecial{"Erro-1"}
	erroComoParametro(testandoErro)

}
