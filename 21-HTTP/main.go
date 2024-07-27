package main

import (
	"log"
	"net/http"
)

//STATUS CODES = Status da requisição, sucesso, problema, not found e tals
//no caso do GO nao precisa lembrar a numeração, passe http. e selecione a que precisa
//w.WriteHeader(http.StatusCreated)

// estrutura para fazer uma função de requisicao
// o func é o request = requisição e o response = resposta, a estrutura já é padronizada
func home(w http.ResponseWriter, r *http.Request) {

	//w.WriteHeader(http.StatusCreated)

	w.Write([]byte("ola usuarios")) // escreve uma mensagem e precisa esta no slice de byte
}
func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ola usuarios bons")) // escreve uma mensagem e precisa esta no slice de byte
}
func main() {
	//HTTP É UM PROTOCOLO DE COMUNICAÇÃO- BASE DE COMUNICAÇÃO WEB
	// CLIENTE (FAZ REQUISIÇÃO) - SERVIDOR (PROCESSA REQUISIÇÃO E ENVIA RESPOSTA)
	// REQUEST - RESPONSE

	// ROTAS
	// URI = IDENTIFICADOR DO RECURSO
	// MÉTODOS = GET, POST, PUT, DELETE

	//estrutura para fazer uma requisição
	//recebe dois parametros, um é a URI (caminho), e o outro é a função
	// o func é o request = requisição e o response = resposta, a estrutura já é padronizada, nesse caso a função esta lá em cima e chamou o metodo
	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", usuarios)

	//estrutura para rodar um servidor
	log.Fatal(http.ListenAndServe(":8080", nil))

}
