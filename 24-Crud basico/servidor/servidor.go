package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id`
	Nome  string `json:"nome`
	Email string `json:"email`
}

// Criar usuario insere o usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("falha ao ler a requisição"))
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("falha ao converter o usuario para struct"))
		return
	}
	fmt.Println(usuario)

	db, erro := banco.Conectar() // abrir o banco

	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco"))
		return
	}
	defer db.Close() // fechando o banco

	// prepare statement =  criar comando de inserção para evitar ataque sql injection
	statement, erro := db.Prepare("insert into usuarios (nome,email) values (?,?)")

	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	//esse função serve para substituir os valores que estava com ? ? na anterior e manter a ordem certa (nome,email)
	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}

	//essa função retorna o id do usuario inserido = geralmente quando o usuario é inserido, retornar o id dele é um costume
	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o id inserido"))
		return
	}

	//STATUS CODES = Status da requisição, sucesso, problema, not found e tals
	//no caso do GO nao precisa lembrar a numeração, passe http. e selecione a que precisa
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("usuario inserido com sucesso! Id: %d", idInserido)))

}

// Busca todos os usuarios do banco
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco"))
		return
	}
	defer db.Close()

	// query = retorna linhas da tabela
	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao buscar usuarios"))
		return
	}
	defer linhas.Close()

	var usuarios []usuario

	//next = para cada linha que for retornada irá executar uma interação
	for linhas.Next() {
		var usuario usuario // criando um usuario do tipo usuario

		//scanneia cada informação da linha em ordem e coloca nas variaveis do usuario
		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao scanear o usuario"))
			return
		}
		usuarios = append(usuarios, usuario)
		w.WriteHeader(http.StatusOK)

		// estrutura para fazer conversão de json e retornar a resposta via http
		if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
			w.Write([]byte("Erro ao converter usuario para json"))
			return
		}

	}
}

// Busca um usuario no banco
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parametro para inteiro"))
		return
	}

	db, erro := banco.Conectar() // abrir o banco

	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco"))
		return
	}
	//defer db.Close() // fechando o banco

	linha, erro := db.Query("select * from usuarios where id =?", ID)
	if erro != nil {
		w.Write([]byte("Erro ao buscar usuario"))
		return
	}
	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao scannear usuario"))
			return
		}
	}
	// estrutura para fazer conversão de json e retornar a resposta via http
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Erro ao converter usuario para json"))
		return
	}
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição!"))
		return
	}

	var usuario usuario
	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.Write([]byte("Erro ao atualizar o usuário!"))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

// DeletarUsuario remove um usuário do banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Erro ao deletar o usuário!"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
