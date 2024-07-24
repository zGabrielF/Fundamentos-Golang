package enderecos

import "testing"

//esta sendo usado para teste unitario
type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

	//t.Parallel() = vai rodar em paralelo com outras funções que também tenha o parallel

	//testando varios cenarios
	cenariosDeteste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida ABC", "Avenida"},
		{"Rodovia dos imigrantes", "Rodovia"},
		{"Praça das rosas", "Tipo Inválido"},
		{"Estrada qualquer", "Estrada"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeteste {
		TipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if TipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				TipoDeEnderecoRecebido,
				cenario.retornoEsperado,
			)
		}
	}
}

// terminal
// go test ./... = vai ir em todos os pacotes do projeto e rodar, caso apareça (cached) = não houve alteração no codigo de teste, por isso ele nao roda novamente para nao perder tempo
// go test -v = modo verboso, vai escrever informações na tela, qual função esta rodando e o tempo
// go test --cover = vai ver quantos % da função ou cenario de modo geral esta sendo coberto
// go test --coverprofile + nome do arquivo .txt , exemplo coverprofile cobertura.txt = irá criar um arquivo txt com algumas informações e para ler arquivo faça
// go tool cover --func= nome do arquivo, nesse caso --func=cobertura.txt = irá mostrar as funções do pacote testada e a %
// go tool cover --html= nome do arquivo, nesse caso --html=cobertura.txt = irá criar uma pagina html mostrando as linhas que nao estao cobertas

//Teste de unidade = testa uma pequena parte do codigo
//para rodar os testes o arquivo precisa ter o _test.go pois o go só reconhece assim
//e para rodar no terminal digite go test
//essa é a estrutura para a função teste TestT sempre começa com Test e uma maiuscula

// func TestTipoDeEndereco(t *testing.T) {

// 	enderecoParaTeste := "Rua ABC"
// 	tipoDeEnderecoEsperado := "Avenida"
// 	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

// 	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
// 		t.Errorf("O tipo recebido é diferente do esperado, esperava %s e recebeu %s",
// 			tipoDeEnderecoEsperado,
// 			tipoDeEnderecoRecebido,
// 		)
// 	}

// }
