package main

import "fmt"

var n int

//Init faz com que a função seja executado primeiro
func init() {
	fmt.Println("Executando o init")
	n = 10
}
func main() {
	fmt.Println("executando main")
	fmt.Println(n)
}
