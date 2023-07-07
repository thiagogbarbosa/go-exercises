package main

import "fmt"

/*- Esse exercício vai reforçar seus conhecimentos sobre conjuntos de métodos.
  - Crie um tipo para um struct chamado "pessoa"
  - Crie um método "falar" para este tipo que tenha um receiver ponteiro (*pessoa)
  - Crie uma interface, "humanos", que seja implementada por tipos com o método "falar"
  - Crie uma função "dizerAlgumaCoisa" cujo parâmetro seja do tipo "humanos" e que chame o método "falar"
  - Demonstre no seu código:
      - Que você pode utilizar um valor do tipo "*pessoa" na função "dizerAlgumaCoisa"
      - Que você não pode utilizar um valor do tipo "pessoa" na função "dizerAlgumaCoisa"  */

type pessoa struct {
	nome  string
	idade int
}

func (p *pessoa) falar() {
	fmt.Println(p.nome, "diz: Olá")
}

type humanos interface {
	falar()
}

func dizerAlgumaCoisa(h humanos) {
	h.falar()
}

func main() {
	pes1 := pessoa{
		nome:  "Thiago",
		idade: 23,
	}

	dizerAlgumaCoisa(&pes1)
}
