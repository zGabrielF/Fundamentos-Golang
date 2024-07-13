package main

import "fmt"

// nesse caso esta passando uma copia de um valor para essa função, usada para manter o valor original da variavel
func inverterSinal(numero int) int {
	return numero * -1
}

// nesse caso esta passando uma referencia de um valor para essa função, usada para mudar o valor original variavel para o codigo inteiro
func inverterComPonteiro(numero *int) {
	*numero = *numero * -1
}
func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
