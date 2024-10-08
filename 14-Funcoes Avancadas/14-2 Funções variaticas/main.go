package main

import "fmt"

// a função recebe de 1 até N numeros int
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

// nesse caso passando dois tipos, não pode ter dois variaticos, ou seja não pode ter duas vezes o ...
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	total := soma(1, 2, 3, 4)
	fmt.Println(total)

	escrever("olá mundo", 10, 2, 3, 4, 5, 6, 7)

}
