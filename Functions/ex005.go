package main

import (
	"fmt"
	"math"
)

type formaGeometrica interface {
	getNome() string
	area() float64
}

func info(f formaGeometrica) {
	fmt.Println("--------------------")
	fmt.Println("Calculando a área da forma geométrica: ", f.getNome())
	fmt.Println("E a área é: ", f.area())
}

type circulo struct {
	nome string
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func (c circulo) getNome() string {
	return c.nome
}

type quadrado struct {
	nome string
	lado float64
}

func (q quadrado) area() float64 {
	return math.Pow(q.lado, 2)
}

func (q quadrado) getNome() string {
	return q.nome
}

func main() {
	c1 := circulo{
		nome: "Circulo",
		raio: 5,
	}
	q1 := quadrado{
		nome: "Quadrado",
		lado: 5,
	}
	info(c1)
	info(q1)
}
