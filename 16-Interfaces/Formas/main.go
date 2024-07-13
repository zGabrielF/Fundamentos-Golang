package main

import (
	"fmt"
	"math"
)

// Interface
// nesse caso para utilizar a interface as funções precisam implementar o metodo area obrigatoriamente
type forma interface {
	area() float64
}

// nessa função esta recebendo uma forma que já esta implementado a interface
func escreverArea(f forma) {
	fmt.Printf("A area da forma é %0.2f\n", f.area())
}

// nessa função esta calculando a area do retangulo
func (r retangulo) area() float64 {
	return r.altura * r.largura
}

// nessa função esta calculando a area do circulo
func (c circulo) area() float64 {
	// o math é um pacote matematico que já contem algumas funções, o pi retorna o valor do pi, e o pow eleva o numero, por exemplo raio elevado a 2
	return math.Pi * math.Pow(c.raio, 2)
}

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

func main() {
	r := retangulo{10, 15}
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)
}
