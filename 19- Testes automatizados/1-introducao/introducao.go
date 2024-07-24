package main

import (
	"fmt"
	enderecos "introducao-testes/2-enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Avenida Paulista")
	fmt.Println(tipoEndereco)
}
