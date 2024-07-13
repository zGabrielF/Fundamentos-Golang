package main

import "fmt"

//retorno nomeado
func calculos(n1, n2 int) (soma, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}
func main() {

	soma, subtracao := calculos(10, 30)
	fmt.Println(soma, subtracao)
}
