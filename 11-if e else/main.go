package main

import "fmt"

func main() {
	fmt.Println(1)

	numero := 20
	if numero > 15 {
		fmt.Println("numero é maior que 15")
	} else {
		fmt.Println("numero é menor que 15")
	}

	// nesse caso estou usando o if e declarando uma variavel, passando outra e testando a condição
	if numero2 := numero; numero2 > 15 {
		fmt.Println("é maior que 15")
	} else {
		fmt.Println("não é maior que 15")
	}
}
