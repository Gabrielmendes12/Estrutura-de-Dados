package lista

import (
	"encoding/json"
	"fmt"
	"net/http"
	"trabap1/dados"
)

func Adicionar(w http.ResponseWriter, r *http.Request){
	mensagem := r.FormValue("mensagem")

	dados.Dados = append(dados.Dados, mensagem)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Write([]byte(mensagem))
	fmt.Println("Mensagem lida e adicionada:", mensagem)
}

func Obter(w http.ResponseWriter, r *http.Request){
	mensagensJSON, err := json.Marshal(dados.Dados)

	if err != nil {
		http.Error(w, "Erro ao converter mensagens para JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(mensagensJSON)
	fmt.Println("Mensagens carregadas com sucesso!")
}

