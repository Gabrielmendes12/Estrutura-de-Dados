package main

import (
	"fmt"
	"net/http"
	"trabap1/handlers"
	"trabap1/handlers/lista"

	"github.com/gorilla/mux"
)

func main(){
	// router
	r := mux.NewRouter()

	// rotas
	r.HandleFunc("/olamundo", handlers.OlaMundo).Methods("GET")
	r.HandleFunc("/lista", lista.Adicionar).Methods("POST")

	fmt.Println("Servidor iniciado em http://localhost:8000")
	http.ListenAndServe(":8000", r)
}