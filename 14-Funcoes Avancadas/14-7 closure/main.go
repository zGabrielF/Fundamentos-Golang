package main

import "fmt"

// funcao que referencia variaveis que estão fora do seu corpo
func closure() func() {
	texto := "Dentro da função closure"

	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}
func main() {
	texto := "dentro do main"
	fmt.Println(texto)

	funcaoClosure := closure()
	funcaoClosure()
}
