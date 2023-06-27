package main

import "fmt"

/*Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
    - "Nome", "Sobrenome", "Hobby favorito"
- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.*/

func main() {

	nameAndHobby := [][]string{
		[]string{"Anderson", "Silva", "Lutar"},
		[]string{"Airton", "Senna", "Correr"},
		[]string{"Marie", "Curie", "Descobrir coisas"},
	}
	for _, v := range nameAndHobby {
		fmt.Println(v[0])
		for _, desc := range v {
			fmt.Println("\t", desc)
		}
	}

}
