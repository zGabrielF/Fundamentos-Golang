package main

import "fmt"

func main() {

	//PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA

	var vari int
	var ponteiro *int // o * indica que ele é um ponteiro, nesse caso um ponteiro do tipo int

	fmt.Println(vari, ponteiro)

	vari = 100
	ponteiro = &vari // o & indica o endereço de memoria onde esta alocado o valor da variavel

	fmt.Println(vari, ponteiro) // aqui ele irá mostrar o valor e o endereço de memoria onde está o valor da variavel

	//nesse caso o * esta fazendo uma desferenciação,ele irá mostrar o valor que está alocado na memoria e não o endereço de memoria
	fmt.Println(vari, ponteiro)

	vari = 150
	fmt.Println(vari, ponteiro) // nesse caso ele foi atualizado o valor e pega o ultimo valor adicionado, permanecendo no mesmo endereço de memoria

}
