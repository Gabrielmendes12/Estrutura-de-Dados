package main

import "fmt"

type Aluno struct {
	Nome string
	Matricula string
	Curso string
	Disciplinas []string
}

func main() {
	al1 := Aluno {
		Nome: "Gabriel",
		Matricula: "1234",
		Curso: "ADS",
		Disciplinas: []string{"Estruturas de Dados", "Desenvolvimento Web"},
	}

	fmt.Println(al1)
	fmt.Println(al1.Nome)
}