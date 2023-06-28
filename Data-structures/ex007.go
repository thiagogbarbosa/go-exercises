package main

import "fmt"

/*Crie um map com key tipo string e value tipo []string.
    - Key deve conter nomes no formato sobrenome_nome
    - Value deve conter os hobbies favoritos da pessoa
- Demonstre todos esses valores e seus índices.*/

func main() {
	nameAndHobby := map[string][]string{
		"silva_Andersson": []string{"To fight", "To dance", "Spend time with family"},
		"senna_Airton":    []string{"To drive", "to Watch movies", "to Code"},

		"curie_Marie": []string{"To discover cool things", "To read", "To make science"},
	}

	//To add a new value on maps

	nameAndHobby["barbosa_Thiago"] = []string{"To code", "To play soccer", "To spend time with my cat"}

	//To remove any value from maps

	delete(nameAndHobby, "silva_Andersson")

	for k, v := range nameAndHobby {
		fmt.Println(k)
		for j, h := range v {
			fmt.Println("\t", j, "-", h)
		}
	}

}
