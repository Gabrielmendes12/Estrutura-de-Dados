package main

import "fmt"

const MAX_ARRAY = 5

// pesquisar características de arrays
func main () {
	var numeros [5]int
	var nomes [MAX_ARRAY]string

	fmt.Println(numeros)
	fmt.Println(nomes)

	numeros[2] = 5
	fmt.Println(numeros)

	//
	outrosNumeros := [4]int{4, 8, 1, -5}
	fmt.Println(outrosNumeros)

	// percorrendo arrays
	for i := 0; i < len(numeros); i++ {
		fmt.Println(i, numero)
	}
}