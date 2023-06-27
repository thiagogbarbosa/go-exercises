package main

import "fmt"

/*Crie uma slice usando make que possa conter todos os estados do Brasil.
"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso",
 "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro",
"Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"
Demonstre o len e cap da slice.
Demonstre todos os valores da slice sem utilizar range.*/

func main() {

	estadosBr := make([]string, 27, 27)
	estadosBr = []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso",
		"Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro",
		"Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}

	fmt.Println(len(estadosBr), cap(estadosBr))

	for i := 0; i < len(estadosBr); i++ {
		fmt.Println(estadosBr[i], i)
	}

}
