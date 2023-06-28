package main

import "fmt"

/*- Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
    - Nome
    - Sobrenome
    - Sabores favoritos de sorvete
- Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.*/

type person struct {
	name                    string
	surname                 string
	FavoriteIceCreamFlavors []string
}

func main() {
	client1 := person{
		name:                    "Marcus",
		surname:                 "Aurelius",
		FavoriteIceCreamFlavors: []string{"Chocolate", "Strawberry", "Mint"},
	}

	client2 := person{
		name:                    "Isaac",
		surname:                 "Newton",
		FavoriteIceCreamFlavors: []string{"Vanilla", "Cookies and Cream", "Pistachio"},
	}

	fmt.Println(client1.name, client1.surname)
	for i, v := range client1.FavoriteIceCreamFlavors {
		fmt.Println("\t", i, "-", v)

	}
	fmt.Println(client2.name, client2.surname)
	for i, v := range client2.FavoriteIceCreamFlavors {
		fmt.Println("\t", i, "-", v)
	}

}
