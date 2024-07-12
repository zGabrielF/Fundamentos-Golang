package main

import "fmt"

// struct recebendo outro struct
type usuario struct {
	nome     string
	idade    uint8
	endereco endereco //colocando um struct dentro de outro, la ele
}

// struct normal
type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("arquivo")

	//1 forma de usar o struct
	var u usuario
	u.nome = "biel"
	u.idade = 20
	println(u.nome, u.idade)

	endereco := endereco{"Rua do riachao", 173}

	// 2 forma de usar o struct
	usu := usuario{"biel", 20, endereco}
	println(usu)

	//3 forma de usar o struct, passando apenas 1 parametro, pode ser o nome, ou a idade e assim por diante
	usu2 := usuario{nome: "biel"}
	println(usu2.nome)

}
