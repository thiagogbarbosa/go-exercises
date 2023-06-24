package main

import "fmt"

/* - Atribua um valor int a uma variável
   - Demonstre este valor em decimal, binário e hexadecimal
   - Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
   - Demonstre esta outra variável em decimal, binário e hexadecimal */

func main() {

	//Step 1

	numb1 := 100
	fmt.Printf("The number %v in decimal is: %d\n", numb1, numb1)
	fmt.Printf("The number %v in binary is: %b\n", numb1, numb1)
	fmt.Printf("The number %v in hexadecimal is: %#x\n", numb1, numb1)

	//Step 2 Bit shifting

	numb1Bitshift := numb1 << 1
	fmt.Printf("The number %v in decimal is: %d\n", numb1Bitshift, numb1Bitshift)
	fmt.Printf("The number %v in binary is: %b\n", numb1Bitshift, numb1Bitshift)
	fmt.Printf("The number %v in hexadecimal is: %#x\n", numb1Bitshift, numb1Bitshift)

}
