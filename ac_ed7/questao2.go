package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	var L, Q float64

	fmt.Scanln(&N, &L, &Q)

	participantes := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&participantes[i])
	}

	cuiaCheia := int(math.Floor(L / Q)) // arredondamento para baixo para garantir somente o uso de cuias completas
	participanteRico := participantes[cuiaCheia % N]

	aguaRestante := math.Mod(L, Q) // resto da divisão da capacidade de água térmica pelo volume da cuia

	fmt.Printf("%s %.1f\n", participanteRico, aguaRestante)
}
