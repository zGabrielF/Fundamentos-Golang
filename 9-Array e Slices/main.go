package main

import (
	"fmt"
)

func main() {

	fmt.Println(1)

	// 1 maneira de declarar o array
	var array [5]string
	array[0] = "posição 1"
	fmt.Println(array)

	// 2 maneira de declarar o array
	array2 := [5]string{"posi 1", "posi 2", "posi 3", "posi 4", "posi 5"}
	fmt.Println(array2)

	// 3 maneira de declarar o array, nesse caso o [...] vai pegar a quantidade de valores passados no {} para ser seu tamanho
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	//o slice é uma fatia de um array, não tem tamanho fixo
	slice := []int{10, 11, 12, 13, 14, 15}
	fmt.Println(slice)

	//função append adiciona um valor novo no slice
	slice = append(slice, 20)
	fmt.Println(slice)

	//esse slice é um pedaço do outro slice também pode ser feito com array, ele pega o indice 1 até o indice 3
	//no caso ele ignora o indice 3 e pega o 1/2

	slice2 := array2[1:3]
	fmt.Println(slice2)

	//slice2 := slice[1:3]
	//fmt.Println(slice2)

	array2[1] = "posi alterada"
	fmt.Println(slice2)

	//reflect função que vai retornar o tipo da variavel
	//fmt.Println(reflect.TypeOf(slice))
	//fmt.Println(reflect.TypeOf(array3))

	//arrays internos

	//função make = alocar um espaço na memoria para uma determinada coisa
	// 3 parametros, tipo, tamanho e a quantidade maxima

	//cria um array de 11 posições e retorna um slice com as 10 primeiras posições, caso você tente estourar a capacidade ele vai dobrar
	//a capacidade atual, sendo assim não permitido estourar

	slice3 := make([]float32, 10, 11)
	slice3 = append(slice3, 20) // adicionado valor para bater a capacidade maxima
	slice3 = append(slice3, 21) // adicionado valor para bater a capacidade maxima
	fmt.Println(slice3)
	fmt.Println((len(slice3))) // lenght = ver o tamanho do slice
	fmt.Println((cap(slice3))) // capacidade = ver a capacidade do array

	// nesse caso não passou a capacidade maxima, ele irá ficar com a capacidade igual ao tamanho, pois é algo opcional
	slice4 := make([]float32, 5)
	slice4 = append(slice4, 20) // adicionado valor para bater a capacidade maxima
	slice4 = append(slice4, 21) // adicionado valor para bater a capacidade maxima
	fmt.Println(slice4)
	fmt.Println((len(slice4))) // lenght = ver o tamanho do slice
	fmt.Println((cap(slice4))) // capacidade = ver a capacidade do array

}
