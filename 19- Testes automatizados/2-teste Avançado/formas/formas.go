package formas

import (
	"math"
)

// Interface
// nesse caso para utilizar a interface as funções precisam implementar o metodo area obrigatoriamente
type Forma interface {
	area() float64
}

type Retangulo struct {
	Altura  float64
	Largura float64
}

type Circulo struct {
	raio float64
}

// nessa função esta calculando a area do retangulo
func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

// nessa função esta calculando a area do circulo
func (c Circulo) Area() float64 {
	// o math é um pacote matematico que já contem algumas funções, o pi retorna o valor do pi, e o pow eleva o numero, por exemplo raio elevado a 2
	return math.Pi * math.Pow(c.raio, 2)
}
