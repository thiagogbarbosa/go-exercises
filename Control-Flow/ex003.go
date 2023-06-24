package main

import "fmt"

/*- Crie um loop utilizando a sintaxe: for condition {}
- Utilize-o para demonstrar os anos desde que você nasceu.*/

func main() {

	birthyear := 2000
	countTo := 2023

	for birthyear < countTo {
		fmt.Println(birthyear)
		birthyear++
	}

}
