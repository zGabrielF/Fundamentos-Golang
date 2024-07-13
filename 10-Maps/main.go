package main

import "fmt"

func main() {

	fmt.Println(1)

	//dentro do [] é o tipo da chave, e fora é o tipo dos valores, e ambos precisam ser iguais, caso seja interface pode usar qualquer tipo
	//no caso for usuario := map[int]string, ou vice e versa ele aceita
	usuario := map[string]string{
		"nome":      "biel",
		"sobrenome": "freitas",
	}
	// nesse caso ele acessa apenas o nome
	//fmt.Println(usuario["nome"])

	fmt.Println(usuario)

	// nesse caso ele esta aninhando um map dentro do outro, é como se fosse uma relação entre tabelas
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Felipe",
			"ultimo":   "santos",
		},
		"curso": {
			"nome":   "engenharia",
			"campus": "Campus1",
		},
	}
	fmt.Println(usuario2)

	delete(usuario2, "nome") //função nativa do go para apagar uma chave, 1 parametro = map , 2 parametro = chave
	fmt.Println(usuario2)

	// para adicionar outra chave dentro do map, precisa seguir a mesma estrutura anterior
	usuario2["signo"] = map[string]string{
		"nome": "Libra",
	}
	fmt.Println(usuario2)
}
