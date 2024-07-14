package main

import (
	"fmt"
	"time"
)

func main() {
	// CONCORRENCIA != PARALELISMO
	// goroutine = ele faz a proxima linha ser executada mesmo que a linha anterior não tenha acabado sua execucao
	go escrever("ola mundo")      // goroutine
	escrever("programando em go") // caso coloque go aqui também não irá imprimir nada, pois não tem mais comando
}
func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
