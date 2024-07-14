package main

import "fmt"

func main() {
	// canal com buffer
	// nessa estrutura, o canal esta recebendo uma capacidade de 2, assim só poderá da deadlock após o 3 valor
	// a diferença entre canal normal e canal com buffer, é que o normal vai da deadlock pois bloqueia, e com capacidade ele só irá bloquear quando passar da capacidade

	canal := make(chan string, 2)
	canal <- "olá"
	canal <- "programando em go"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)

	fmt.Println(mensagem2)
}
