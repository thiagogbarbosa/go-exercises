package main

import "fmt"

/*- Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
- Demonstre estes valores.*/

const (
	_ = 2023 + iota
	nextyear
	next2year
	next3year
	next4year
)

func main() {
	fmt.Println(nextyear, next2year, next3year, next4year)

}
