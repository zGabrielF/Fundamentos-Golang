package main

import (
	"errors"
	"fmt"
)

func main() {
	//tipos de dados inteiro, o int sem numeração vai pegar pela arquitetura, ou seja 64 ou 32 bits
	//int8, int16, int32, int64, int

	//tipos de dados unsygned inteiro, inteiro sem sinal, ou seja nao aceita numero negativo
	//uint8, uint16, uint32, uint64, uint

	// int32 = rune, nomeação (apelido)
	//var numero2 rune =123

	// uint8= byte, nomeação (apelido)
	//var numero3 rune =123

	//tipos de dados float
	//float32, float64

	// tipo de dados booleano, true ou false, se declarar sem o valor sempre será false
	var booleano bool = true
	fmt.Println(booleano)

	// tipo de dados erro, retorna o valor 0 = <nil>
	var erro error
	fmt.Println(erro)

	// tipo de dados erro, retornando uma mensagem de erro
	var erro2 error = errors.New("Erro")
	fmt.Println(erro2)

	var numero int16 = 100

	fmt.Println(numero)
}
