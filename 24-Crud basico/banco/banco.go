package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Driver de conexão com o Mysql
)

// go get github.com/go-sql-driver/mysql = para baixar o driver do mysql
// go mod init + nome do modulo = para criar um modulo

// Conectar abre a conexão com o banco de dados
func Conectar() (*sql.DB, error) {
	urlConexao := "biel2:golang@/devbook?charset=utf8&parseTime=true"

	db, erro := sql.Open("mysql", urlConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}
	return db, nil
}
