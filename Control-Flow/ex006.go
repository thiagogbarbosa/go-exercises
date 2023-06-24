package main

import "fmt"

/*Crie um programa que utilize a declaração switch, sem switch statement especificado.*/

func main() {
	employeeName := "Andreza"

	switch {
	case employeeName == "João" || employeeName == "Maria":
		fmt.Println("Team A is here!")
	case employeeName == "Anderson" || employeeName == "Sophie":
		fmt.Println("Team B is here!")
	case employeeName == "Marcus" || employeeName == "Noah":
		fmt.Println("Team C is here!")
	default:
		fmt.Println("Are you employed?")
	}

}
