package main

import "fmt"

// recover = usado para recuperar a execução caso esteja em panic, essa é a estrutura e precisa usar o defer antes de chamar a função
func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada")
	}
}
func alunoAprovado(n1, n2 float64) bool {
	defer recuperarExecucao() // usando para recuperar a execução que estava em panic
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("a media é 6") // o panic faz com que a execução "morra" interrompendo o fluxo normal e chamando as funções default
}
func main() {
	fmt.Println(alunoAprovado(6, 6))
	fmt.Println("executou")

}
