package main

import "fmt"

/*Usando uma literal composta:
    - Crie uma slice de tipo int
    - Atribua 10 valores a ela
- Utilize range para demonstrar todos estes valores.
- E utilize format printing para demonstrar seu tipo.*/

func main() {

	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//Showing the slice values
	for _, v := range mySlice {
		fmt.Println(v)
	}
	//Printing the format of my slice, as proposed
	fmt.Printf("Type of mySlyce: %T", mySlice)

	//Slicing my slice:
	//Do primeiro ao terceiro item do slice (incluindo o terceiro item!)

	fmt.Println(mySlice[:3])

	//Do quinto ao último item do slice (incluindo o último item!)

	fmt.Println(mySlice[4:])

	//Do segundo ao sétimo item do slice (incluindo o sétimo item!)

	fmt.Println(mySlice[1:7])

	//Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)

	fmt.Println(mySlice[3:9])

	//Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item

	fmt.Println(mySlice[3 : len(mySlice)-1])

}
