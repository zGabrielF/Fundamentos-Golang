package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("olá")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}

}

//Generator = vai ser uma função que vai encapsular uma chamada para uma goroutine e devolver um canal

// essa função vai retornar um canal de string de apenas uma direção, ele só recebe
func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return canal
}
