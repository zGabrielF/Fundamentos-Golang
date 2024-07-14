package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplexar(escrever("olá"), escrever("programando em go"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

// Multiplexador = serve para pegar dois canais e juntar em apenas um
// essa função
func multiplexar(canalDeEntrada1, canalDeEntrada2 <-chan string) <-chan string {
	canalDeSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canalDeEntrada1:
				canalDeSaida <- mensagem
			case mensagem := <-canalDeEntrada2:
				canalDeSaida <- mensagem
			}
		}
	}()
	return canalDeSaida
}

// essa função vai retornar um canal de string de apenas uma direção, ele só recebe
func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
			//o time duration calcula o tempo da execução e o rand.intn é para gerar um numero aleatorio de 0 até 2000
		}
	}()
	return canal
}
