package main

import "fmt"

func main() {

	//Função Anonima é uma função que não precisa de nome

	// nesse caso como não tem parametro, o valor vai dentro função
	func() {
		fmt.Println("Olá")
	}()

	// nesse caso passando o parametro o valor precisa ser colocado lá no fim
	func(texto string) {
		fmt.Println(texto)
	}("Olá")

	// nesse caso passando a função para uma variavel, pois sem a variavel não irá imprimir a informação
	retorno := func(texto string) string {
		return fmt.Sprintf("recebido -> %s", texto) // usando para concatenar informações %s = string , %d = numero
	}("Olá")
	fmt.Println(retorno)
}
