package main

func main() {
	println(1)

	i := 0

	//Repetições
	//Enquanto
	for i < 10 {
		i++
		println("valor: ", i)
		//time.Sleep(time.Second) // faz com que o codigo atrase um pouco a cada segundo
	}

	//Até caso precise pular em mais de um, colocar j +=2 por exemplo
	for j := 0; j < 10; j++ {
		println("valor: ", j)
		//time.Sleep(time.Second) // faz com que o codigo atrase um pouco a cada segundo
	}

	//RANGE

	//range quando precisa percorrer pelo os valores
	nomes := [3]string{"João", "Biel", "Lucas"}

	//for com range, essa é estrutura, sempre indice e depois o valor do item, caso precise ignorar um coloque _ no lugar
	for indice, nome := range nomes {
		println(indice, nome)
	}
	// nesse caso ele pega o valor da tabela ascii, caso precise pegar o valor, string(letra)
	for indice, letra := range "Palavra" {
		println(indice, letra)
	}

	//caso precise pegar o valor, string(letra)

	/*for indice, letra := range "Palavra" {
		println(indice, string(letra))
	}*/

	// range com map
	usuario := map[string]string{
		"nome":      "Biel",
		"sobrenome": "Freitas",
	}
	for chave, valor := range usuario {
		println(chave, ":", valor)
	}
}
