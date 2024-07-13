package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

// essa função esta salvando o dado do usuario
func (u usuario) salvar() {
	fmt.Printf("Salvando dados %s \n", u.nome)
}

// essa função verifica se a idade é maior que 18
func (u usuario) maiorIdade() bool {
	return u.idade >= 18
}

// essa função esta referenciando pelo ponteiro assim alterando o valor da idade
func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {

	// crie objeto com nome diferente de usuario, para nao ter erro de compilação por exemplo usuario = usuario 1; e usuario2;
	usuario1 := usuario{"biel", 20}
	usuario1.salvar()

	usuario2 := usuario{"davi", 20}
	usuario2.salvar()

	maiorDeIdade := usuario2.maiorIdade()
	fmt.Println(maiorDeIdade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)

}
