package main

import "fmt"

func somar(n1 int, n2 int) int {
	return n1 + n2
}

/* nessa função os parametros pode ser escrito dessa forma, devido ambos serem tipo int, ou seja, não precisa botar int em ambos, caso seja tipos
diferente ai precisa

e ela também permite dois retornos ou mais em uma mesma função, o maximo é até 4, algo especifico do GO
*/

func calculos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {

	soma := somar(5, 10)
	fmt.Println("resultado soma", soma)

	/*
		nesse caso ficará dando erro no calculos(10,20), devido ter dois retornos e aqui está recebendo apenas um
		resultadoCalculo := calculos(10, 20)
		fmt.Println(resultadoCalculo)
	*/

	// forma correta de fazer para os dois retornos, e também dando os dois retornos
	resultadoSoma, resultadoSubtracao := calculos(10, 20)
	fmt.Println("resultado soma", resultadoSoma, "resultadoSubstracao", resultadoSubtracao)

	/* nesse caso como tem dois retornos, e você só precise de um resultado ele pode ser ignorando, colocando _ no lugar do que será ignorado
	resultadoSoma, _ := calculos(10, 20)
	fmt.Println("resultado soma", resultadoSoma)
	*/
}
