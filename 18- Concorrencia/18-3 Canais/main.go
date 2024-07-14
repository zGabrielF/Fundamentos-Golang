package main

import (
	"fmt"
	"time"
)

func main() {
	// canal = canal de comunicação, pode enviar e receber dados e ele tem um bloqueio, assim gerando deadlock
	// essa é a estrutura o chan é palavra chave
	canal := make(chan string)
	go escrever("ola mundo", canal)

	// o canal tarefas só recebe dados, o canal resultados só envia dados
	//func worker(tarefas <-chan int, resultados chan<- int) {

	// deadlock = você não tem nenhum local enviando dados, mas seu canal ainda espera receber um dado
	// OBS = ter cuidado com deadlock, pois não é pego em compilação, apenas em execução

	// fica repetindo até gerar as 5 vezes da função e dará um deadlock
	// para resolver o deadlock, precisa verificar se o canal esta aberto ou fechado, é uma propriedade do canal
	// e colocar close(canal) no fim do for na função

	/* for{
		mensagem := <- canal
		fmt.Println(mensagem)
	}*/

	// <- canal = recebendo um valor, canal esta esperando receber um valor
	// nesse caso passa o , aberto para verificar se o canal esta ou não aberto

	// 1 maneira de fazer
	for mensagem := range canal {
		fmt.Println(mensagem)
	}
	fmt.Println("fim do programa")
	// 2 maneira de fazer
	/*mensagem, aberto := <-canal
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}*/
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // <- = esta mandando um valor para dentro do canal, no caso o valor de texto
		time.Sleep(time.Second)
	}
	close(canal) // fechando o canal
}
