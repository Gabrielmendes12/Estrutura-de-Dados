package main

import "fmt"

// Problema de algorítmo (Tomadas) -> slide 8 da aula

func tomadas (){
	var t1, t2, t3, t4 int8
	fmt.Scanln(&t1, &t2, &t3, &t4)
	fmt.Println(t1 + t2 + t3 + t4 - 3)

}

/* Problema de algorítmo (Meteoros) -> slide 8 da aula

x1 <= x <= x2
y2 <= y <= y1

num = []
j = 0
Enquanto VERDADEIRO faça :
	Ler x1, y1, x2, y2
	Se x1 = y1, x2 = y2 != 0 então PODE
	Ler n

	Para cada i = 1, 2, ..., n faça:
		Ler x, y
		Se x1 <= x e x1 <= x2 e y2 <= y e y <= y1, então:
			num[i] = num[i] + 1
	i = i + 1

	Para cada, j = 0, ..., comprimento de num faça:
		Escrever "Teste " +(j+1)
		Escrever num[j]


*/

func meteoritos (){
	var x1, x2, y1, y2, x, y, n int
	num := make([]int, 0)
	i := 0
	for {
		fmt.Scanln(&x1, &y1, &x2, &y2)
		if x1 == 0 && y1 == 0 && x2 == 0 && y2 == 0 {break}
		fmt.Scanln(&n)
		num = append(num, 0)
		for j := 1; j <= n; j++ {
			fmt.Scanln(&x, &y)
			if x1 < x && x <= x2 && y2 <= y && y <= y1{
				num [i]++
			}
		}
		i++
	}
	for j, num_metoritos := range num {
		fmt.Println("Teste", j + 1)
		fmt.Println(num_metoritos)
		fmt.Println("")
	}
}


// Problema de algorítmo (Maior Número) -> slide 8 da aula

func maiorNumero(){

	var numero int
	maior := 0

	for {
		fmt.Scan(&numero)
		if numero == 0 {break}
		if numero > maior {
			maior = numero
		}
	}
	fmt.Println(maior)
}


// Exercício implementação função -> slide 25 da aula

// Função busca que recebe uma matriz, o tamanho n e o valor a ser procurado x
func buscaMatriz(m [][]int, n, x int) bool {  // m é a matriz
	var i, j int
	i = 0

	for i < n {
		j = 0

		// Loop interno para percorrer as colunas da matriz
		for j < n {
			// Verifica se o elemento na posição atual é igual a x
			if m[i][j] == x {
				return true // Retorna verdadeiro se encontrou x na matriz
			}

			// Incrementa j para ir para a próxima coluna
			j++
		}
		// Incrementa i para ir para a próxima linha
		i++
	}

    // Se chegou aqui, não encontrou x na matriz, então retorna falso
    return false
}


func main (){
	//tomadas()
	// meteoritos()
	// maiorNumero()


	matriz := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// Valor a ser procurado
	valorProcurado := 5

	// Chame a função buscaMatriz e imprima o resultado
	encontrou := buscaMatriz(matriz, len(matriz), valorProcurado)

	if encontrou {
		fmt.Printf("O valor %d foi encontrado na matriz.\n", valorProcurado)
	} else {
		fmt.Printf("O valor %d não foi encontrado na matriz.\n", valorProcurado)
	}
}


// Exercício algoritmo em pseudocódigo -> slide 32 de aula




