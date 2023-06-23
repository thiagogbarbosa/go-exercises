package main

import "fmt"

/*- Crie um tipo. O tipo subjacente deve ser int.
- Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.
- Na função main:
    1. Demonstre o valor da variável "x"
    2. Demonstre o tipo da variável "x"
    3. Atribua 42 à variável "x" utilizando o operador "="
    4. Demonstre o valor da variável "x"
*/
type newtype int

var x newtype

//Para testar a conversão no tipo subjacente:

var y int

func main() {
	//Demosntre o valor da variável x
	fmt.Println(x)

	//Demonstre o tipo da variável x
	fmt.Printf("Type of x is: %T\n", x)

	//Atribuindo e demonstrando o valor de x
	x = 42
	fmt.Println(x)

	//y como tipo subjacente de x
	y = int(x)
	fmt.Printf("Type of y is: %T\n", y)
	fmt.Println(y)

}
