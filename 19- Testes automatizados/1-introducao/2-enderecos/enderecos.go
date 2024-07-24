package enderecos

import (
	"strings"
)

// TipoDeEndereco essa função verifica se um endereço é válido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposEnderecosValidos := []string{"Rua", "Avenida", "Estrada", "Rodovia"}

	primeiraPalavraDoEndereco := strings.Split(endereco, " ")[0]

	enderecoTemTipoValido := false

	for _, tipo := range tiposEnderecosValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemTipoValido = true
		}
	}
	if enderecoTemTipoValido {
		return primeiraPalavraDoEndereco
	}
	return "Tipo Inválido"
}
