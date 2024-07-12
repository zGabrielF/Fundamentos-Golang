package main

import "fmt"

type pessoa struct {
	nome  string
	idade uint8
}

type estudante struct {
	//pessoa *pessoa nesse caso é usando ponteiro

	pessoa
	curso string
}

func main() {

	//usando ponteiros
	/*p1 := pessoa{"Biel", 20}
	e1 := estudante{&p1, "bsi"}
	fmt.Println("primeiro print", e1)

	//e1.pessoa = &pessoa{"gustavo", 15}
	fmt.Println("segundo print", e1)

	e1.pessoa.nome = "batata"

	fmt.Println(p1)
	fmt.Println("terceiro print", e1)
	*/

	p1 := pessoa{"Biel", 20}
	e1 := estudante{p1, "bsi"}
	fmt.Println("primeiro print", e1)

	//e1.pessoa = &pessoa{"gustavo", 15}
	fmt.Println("segundo print", e1)

	e1.pessoa.nome = "batata"

	fmt.Println(p1)
	fmt.Println("terceiro print", e1)

}
