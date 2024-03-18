package main

import (
	"fmt"
)

func main() {
	var n int
	for {
		fmt.Println("Entre com a quantidade de cartas: ")
		fmt.Scanln(&n)
		if n == 0 {
			break
		}
		descartadas, restante := simularJogo(n)
		fmt.Println("Discarded Cards:", descartadas)
		fmt.Println("Remaining Card:", restante)
	}
}



func simularJogo(n int) (descartadas string, restante int) {
	var cartasDescartadas []int
	for i := 1; i <= n; i++ {
		cartasDescartadas = append(cartasDescartadas, i)
	}
	for len(cartasDescartadas) >= 2 {
		descartadas = descartadas + fmt.Sprintf("%d, ", cartasDescartadas[0])
		cartasDescartadas = cartasDescartadas[1:]
		cartasDescartadas = append(cartasDescartadas, cartasDescartadas[0])
		cartasDescartadas = cartasDescartadas[1:]
	}
	restante = cartasDescartadas[0]
	descartadas = descartadas[:len(descartadas)-2] // remove a vírgula e o espaço extra no final
	return descartadas, restante
}
