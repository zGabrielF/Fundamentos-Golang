package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500) // multiplica por 500 milisegundos
			canal1 <- "canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2) // multiplica por 2 segundos
			canal1 <- "canal 2"
		}
	}()

	for {
		// o select é mais utilizado no caso de concorrencia
		// nesse caso se as mensagem então chegando no canal ele já imprime, sem esperar o outro
		// pois na função acima, um canal recebe em meio segundo, e o outro em 2 segundos, isso faria que um esperasse o outro executar
		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)
		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}
	}
}
