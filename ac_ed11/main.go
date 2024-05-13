package main

import "fmt"

// Questão 1
func distancia(){

	var distancia int;
	fmt.Scanln(&distancia)

	tempo := distancia * 2; // afastamento ocorre de 1 km a cada 2 minutos
	fmt.Printf("%d minutos\n", tempo);
}

// Questão 2
func feira(){
	var N int // quantidade de idas à feira
	fmt.Scanln(&N)

	for i := 0; i < N; i++ {
		var M int // quantidade de produtos disponíveis
		fmt.Scanln(&M)

		produtos := make(map[string]float64) // mapa para armazenar os preços dos produtos
		for j := 0; j < M; j++ {
			var nome string
			var preco float64
			fmt.Scan(&nome, &preco)
			produtos[nome] = preco
		}

		var P int // quantidade de produtos que Dona Parcinova deseja comprar
		fmt.Scanln(&P)

		total := 0.0 // total a ser gasto por Dona Parcinova
		for j := 0; j < P; j++ {
			var nome string
			var quantidade int
			fmt.Scan(&nome, &quantidade)
			total += produtos[nome] * float64(quantidade)
		}

		fmt.Printf("R$ %.2f\n", total)
	}
}

// Questão 3
func avioes_papel(){
	var competidores, folhasCompradas, folhasPorCompetidor int
    fmt.Scanln(&competidores, &folhasCompradas, &folhasPorCompetidor)

    folhasNecessarias := competidores * folhasPorCompetidor

    if folhasNecessarias <= folhasCompradas {
        fmt.Println("S")
    } else {
        fmt.Println("N")
    }
}

// Questão 4
func sequencia() {
	var N int
	fmt.Scanln(&N)

	numeros := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanln(&numeros[i])
	}

	marcados := 0
	for i := 1; i < N-1; i++ {
		if numeros[i] == 2 && numeros[i-1] == 1 && numeros[i+1] == 1 {
			marcados++
			i++ // Pular o próximo número pois ele não pode ser marcado
		}
	}

	fmt.Println(marcados)
}


func main(){
	distancia();
	feira();
	avioes_papel();
	sequencia();


}