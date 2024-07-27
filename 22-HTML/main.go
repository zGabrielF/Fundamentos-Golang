package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {

	templates := template.Must(template.ParseGlob("*.html*"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		u := usuario{"Biel", "Biel2@gmail.com"}
		templates.ExecuteTemplate(w, "home.html", u) // aqui passa o response, o nome do template e algum dado para passar
	})

	//estrutura para rodar um servidor
	fmt.Println("executando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
