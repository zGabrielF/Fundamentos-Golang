package main

import "fmt"

// o defer faz adiar a execução até o ultimo momento possivel, ou seja irá passar por tudo e depois chegará a ele, bem utilizado para banco de dados
func alunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("Media calculada") // defer = adiar a execução
	fmt.Println("Passou pela função aluno aprovado")
	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false
}
func main() {
	fmt.Println(alunoAprovado(7, 8))

}
