package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"` //quando for exportar para json, sempre os atributos em maiusculo
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	// Marshal = converter um map ou um struct para json
	//json.Marshal()
	// Unmarshal = converter um json para map ou um struct
	//json.Unmarshal()

	c := cachorro{"Rex", "Dalmata", 3}
	fmt.Println(c)

	CachorroEmJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(CachorroEmJSON) //aqui ele irá trazer um slice de numeros

	fmt.Println(bytes.NewBuffer(CachorroEmJSON)) // bytes = recebe um slice de btyes e retorna um json

	c2 := map[string]string{
		"nome": "mel",
		"raca": "caramelo",
	}

	Cachorro2EmJSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(Cachorro2EmJSON) //aqui ele irá trazer um slice de numeros

	fmt.Println(bytes.NewBuffer(Cachorro2EmJSON)) // bytes = recebe um slice de btyes e retorna um json

}
