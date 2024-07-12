package main

//import "fmt"

func main() {

	//ARITMETICOS
	// + - / * %

	soma := 1 + 2
	subtracao := 3 - 2
	divisao := 4 / 2
	multiplicacao := 2 * 5
	restoDivisao := 10 % 2

	println(soma, subtracao, divisao, multiplicacao, restoDivisao)
	/* nesse caso dará erro no print devido a ser tipo diferentes de int, 16 e 32, precisa ser ambos iguais
	var num1 int16 = 10
	var num2 int36 = 20
	soma2 := num1 + num2
	println(soma)
	*/

	//ATRIBUIÇÃO

	var vari1 = "Biel"
	vari2 := "Biel2"

	println(vari1, vari2)

	//OPERADORES RELACIONAIS
	//Retorna true ou false

	println(1 > 2)
	println(1 < 2)
	println(1 >= 2)
	println(1 <= 2)
	println(1 != 2)
	println(1 == 2)

	// OPERADORES LOGICOS
	// && || !   E OU NEGAÇÃO

	println("")
	verdadeiro, falso := true, false
	println(verdadeiro && falso)
	println(verdadeiro || falso)
	println(!verdadeiro)
	println(!falso)

	println("")

	//OPERADORES UNARIOS
	numero1 := 10
	numero1++
	numero1 += 15 // numero = numero + 15
	println(numero1)

	numero1--
	numero1 -= 15 // numero = numero - 15

	numero1 *= 15 // numero = numero * 15

	numero1 /= 15 // numero = numero / 15

	numero1 %= 15 // numero = numero % 15

	println(numero1)

	// não tem operador ternario

}
