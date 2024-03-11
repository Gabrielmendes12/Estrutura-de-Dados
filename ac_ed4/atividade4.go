package main

import "fmt"

// Questão 1

// Temos n discos, no problema vamos movimentá-los da torre 'origem' para a torre 'destino' usando a torre 'auxiliar' como torre intermediária.

func torreDeHanoi(n int, origem, auxiliar, destino string){
	if n == 1 { // Se tivermos 1 único disco ele será movido diretamente para torre de destino
		fmt.Printf("Mova o disco 1 da %s para %s\n", origem, destino)

	} else {
		torreDeHanoi(n-1, origem, destino, auxiliar) // Senão, moveremos os n-1 discos da torre de origem para a torre auxiliar usando a torre de destino como auxiliar.

		fmt.Printf("Mova o disco %d da %s para %s\n", n, origem, destino) // Movendo o disco restante da torre de origem para a torre de destino.
		
		torreDeHanoi(n-1, auxiliar, origem, destino) // Agora, ,overemos n-1 discos da torre auxiliar para a torre de destino usando a torre de origem como auxiliar.
	}
}

func main(){

	// Chamada da função da questão 1

	numDiscos := 3 // Número de discos na Torre de Hanói
    origem := "Torre A"
	auxiliar := "Torre B"
	destino := "Torre C"
    
    fmt.Printf("Resolvendo a Torre de Hanói com %d discos:\n", numDiscos)
    torreDeHanoi(numDiscos, origem, auxiliar, destino)


	// Chamada da função da questão 2

	// Exemplo de uso
	 array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	 somaAlvo := 10
 
	 resultado1, resultado2 := ParComSoma(array, somaAlvo)
 
	 if resultado1 == -1 && resultado2 == -1 {
		 fmt.Println("Nenhum par encontrado.")
	 } else {
		 fmt.Printf("Par encontrado: (%d, %d)\n", resultado1, resultado2)
	 }

}

// Questão 2

// A função encontra um par de elementos no array cuja soma é igual ao somaAlvo

func ParComSoma(arr []int, somaAlvo int) (int, int) {
    esquerda, direita := 0, len(arr)-1

	 // Algoritmo de 2 ponteiros para percorrer o array
	for esquerda < direita {
       somaAtual := arr[esquerda] + arr[direita] // a soma será o valor armazenado nos ponteiros

        if somaAtual == somaAlvo{
            return arr[esquerda], arr[direita]
        } else if somaAtual < somaAlvo {
            esquerda++ 
        } else {
            direita--
        }
    }

    // Se nenhum par encontrado retorna os elementos -1, -1
    return -1, -1
}






