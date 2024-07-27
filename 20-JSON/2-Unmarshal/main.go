package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// `json:"-"`, caso bote assim na hora da conversão o campo não sera reconhecido, logo sera ignorado e nao aparece
type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	// Unmarshal = converter um json para map ou um struct
	//json.Unmarshal()

	cachorroEmJson := `{"nome":"Rex","raca":"Dalmata","idade":3}`

	//c:= cachorro{} //objeto vazio

	var c cachorro
	//a estrutura do unmarshal precisa receber um slice de byte dos dados que quer passar e o endereço de memoria do objeto onde ficara os dados
	if erro := json.Unmarshal([]byte(cachorroEmJson), &c); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c)

	// ter cuidado na criação do map, caso o retorno seja diferente do tipo, pois não irá retornar nesse caso
	cachorro2EmJson := `{"nome":"Mel","raca":"Caramelo"}`
	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorro2EmJson), &c2); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c2)

}
