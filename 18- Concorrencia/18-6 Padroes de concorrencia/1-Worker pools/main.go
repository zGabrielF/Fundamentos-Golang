package main

import "fmt"

func main() {

	//canal com buffer
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	//utilizando dois ou mais gouroutines agiliza a execução
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas) // fechando o canal

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}

}

//padrão worker
//Worker pools = é basicamente uma coleção de threads que ficam esperando tarefas serem atribuídas a elas.
// ele é utilizando quando tiver uma fila de tarefas grandes para ser executados, varios processos e cada um pegando um item da fila para executar
//o canal tarefas só recebe dados, o canal resultados só envia dados
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}

}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
