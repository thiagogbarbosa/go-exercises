package main

import "fmt"

/*- Crie e use um struct anônimo.
- Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.*/

func main() {

	anônimo := struct {
		nome      string
		sabor     string
		ondetem   []string
		vaibemcom map[string]string
	}{
		nome:    "Tiramisu",
		sabor:   "Café",
		ondetem: []string{"Shopping", "Supermercado"},
		vaibemcom: map[string]string{
			"Sobremesa":     "Café",
			"Doce da Tarde": "Café também",
		},
	}
	fmt.Println(anônimo)
}
