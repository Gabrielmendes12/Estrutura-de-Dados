package main

import (
	"fmt"
	"strconv"
)

func main() {
	var D int // dígito com falha
	var N string // número original no contrato

	for {
		fmt.Scan(&D, &N)

		// Verifica se é o fim da entrada
		if D == 0 && N == "0" {
			break
		}

		// Substitui o primeiro dígito com falha por 0
		for i := range N {
			if int(N[i]-'0') == D {
				N = N[:i] + "0" + N[i+1:]
				break
			}
		}

		// Converte a string resultante para um inteiro
		resultInt, _ := strconv.Atoi(N)

		fmt.Println(resultInt)
	}
}