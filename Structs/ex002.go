package main

import "fmt"

/*Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
- Demonstre os valores do map utilizando range.
- Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.*/

type person struct {
	name                    string
	surname                 string
	FavoriteIceCreamFlavors []string
}

func main() {

	client1 := map[string]person{
		"barbosa": person{
			name:                    "Thiago",
			surname:                 "Barbosa",
			FavoriteIceCreamFlavors: []string{"Chocolate", "Mint", "Marshmallow"}}}

	for k, v := range client1 {
		fmt.Println(k)
		for j, h := range v.FavoriteIceCreamFlavors {
			fmt.Println("\t", j, "-", h)
		}
	}

}
