package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//wait group = grupo de espera, pacote padrão do GO
	var waitGroup sync.WaitGroup

	//aqui esta passando que tem duas goroutines
	waitGroup.Add(2)

	// 1 goroutine
	go func() {
		escrever("ola mundo")
		waitGroup.Done() // -1 do contador (waitGroup.add(2))
	}()

	// 2 goroutine
	go func() {
		escrever("programando em go")
		waitGroup.Done() // -1 do contador (waitGroup.add(2))
	}()
	waitGroup.Wait() // quando acabar as goroutine ele finaliza
}
func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
