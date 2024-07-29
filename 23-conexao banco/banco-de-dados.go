package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //import explicito
)

func main() {
	// urlConexao := "usuario:senha@/banco" estrutura basica

	urlConexao := "biel2:golang@/devbook?charset=utf8&parseTime=true"
	db, erro := sql.Open("mysql", urlConexao)

	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println("Conexão esta aberta")

	//estrutura para o select
	linhas, erro := db.Query("select * from usuarios")

	if erro != nil {
		log.Fatal(erro)
	}
	defer linhas.Close()

	fmt.Println(linhas)
}
