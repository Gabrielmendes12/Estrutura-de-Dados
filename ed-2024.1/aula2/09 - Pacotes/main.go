package main

import (
	"fmt"
	"projeto/modelos"
	"projeto/modelos/academico"
)


func main(){
	fmt.Println("Olá, mundo")

	aluno := modelos.Aluno{}
	aluno.Nome = "abcd"
	aluno.Matricula = "1234"

	curso := academico.Curso{Nome: "Engenharia"}

	fmt.Println(aluno, curso)
}

