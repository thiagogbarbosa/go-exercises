package main

import "fmt"

/* - Utilize a declaração defer de maneira que demonstre que sua execução só ocorre ao final do contexto ao qual ela pertence. */

func test() {
	defer fmt.Println("Show this latter")
	fmt.Println("Show this first")
}

func main() {
	test()

}
