package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var x, y, z int

	// as 2 linhas abaixo definem a formação do input do usuário no golang
	fmt.Print("Insira um valor inteiro a seguir: ")
	fmt.Scanln(&x)

	fmt.Println(x)

	fmt.Print("Insira dois valores: ")
	fmt.Scanln(&y, &z)
	fmt.Println(y, z)

	var nome string
	fmt.Print("Informe seu nome: ")
	fmt.Scanln(&nome)
	fmt.Println(nome)

	leitor := bufio.NewReader(os.Stdin)
	fmt.Print("Informe seu nome: ")
	nome, err := leitor.ReadString('\n')
	fmt.Println(nome, err)

}

// incompleto