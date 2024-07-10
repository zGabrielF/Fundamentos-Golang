package main

import "fmt"

func main() {

	// tipo da variavel explicitos, o tipo da variavel é declarada juntamente
	var vari1 string = "biel"

	// tipo da variavel inferencia de tipo, o tipo da variavel é atribuida pelo valor que foi dado a ela
	vari2 := "biel 2"

	// outras maneiras de declarar varias variaveis ao invez de "uma por uma"
	var (
		vari3 string = "bb"
		vari4 string = "cc"
	)
	vari5, vari6 := "dd", "ee"

	//Declaração de constante
	const const1 int = 1

	//Fazendo inversão de valor na variavel
	vari5, vari6 = vari6, vari5

	fmt.Println(vari1)
	fmt.Println(vari2)
	fmt.Println(vari3)
	fmt.Println(vari4)
	fmt.Println(vari5)
	fmt.Println(vari6)

}
