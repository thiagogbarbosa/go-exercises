package main

import "fmt"

/*- Crie um tipo struct "pessoa" que contenha os campos:
    - nome
    - sobrenome
    - idade
- Crie um método para "pessoa" que demonstre o nome completo e a idade;
- Crie um valor de tipo "pessoa";
- Utilize o método criado para demonstrar esse valor.*/

type person struct {
	name    string
	surname string
	age     int
}

func (p person) completeName() {
	fmt.Println("The client name is: ", p.name, p.surname)
}
func main() {
	client1 := person{"Thiago", "Barbosa", 23}
	client1.completeName()
}
