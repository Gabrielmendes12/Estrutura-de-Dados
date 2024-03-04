package main

import ("fmt"
"math"
"exercicio4/geometria"
)

// Exercício 1 -> realizado dentro da func main


// Exercício 2

// Função para inverter a sequência de caracteres em uma string
func inverteString(input string) string {
	runes := []rune(input)

	// Inversão dos caracteres no slice de runes usando 2 ponteiros
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {  
		runes[i], runes[j] = runes[j], runes[i]  // índices trocados para executar a função
	}

	// 2 índices (i -> começa no início do slice (0) e j -> começa no final do slice (len(runes-1))
	// O índice i é incrementado em 1 e o índice j é decrementado em 1 para que eles se encontrem no vetor

	// Conversão do slice de runes para um tipo string
	return string(runes)
}

// Exercício 3

// Definindo a estrutura Ponto
type Ponto struct {
	X float64
	Y float64
}

// Método DistanciaOrigem para calcular a distância até a origem (0,0)
func (p *Ponto) DistanciaOrigem() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}
/* Este método é associado ao tipo Ponto através da declaração (p *Ponto). 
Ele calcula a distância até a origem usando a fórmula da distância entre 2 pontos: Raiz de (X^2 + Y^2) -> já que X1, Y1 = (0,0) */


func main(){

	// Exercício 1
	vetor := [10]int{0,1,2,3,4,5,6,7,8,9}

	fmt.Println("Os elementos do vetor são:")
	for i:= 0; i < len(vetor); i++ {  // iteração começa com 0 e vai até o momento onde o item for menor que o tamanho do vetor
		fmt.Println(vetor[i])
	}

	// Chamada da função do exercício 2

	// Entrada do usuário
	fmt.Print("Digite uma palavra: ")
	var palavra string
	fmt.Scanln(&palavra)

	// Chamar a função para inverter a string
	resultado := inverteString(palavra)
	
	fmt.Println("String invertida:", resultado)


	// chamada da função do exercício 3

	// Exemplo de uso
	coordenada := Ponto{X: 3, Y: 4}
	distancia := coordenada.DistanciaOrigem()

	fmt.Printf("Coordenadas do ponto: (%.2f, %.2f)\n", coordenada.X, coordenada.Y)
	fmt.Printf("Distância até a origem: %.2f\n", distancia)

	// chamada da função do exercício 4
	var base, altura float64

    fmt.Print("Digite a base do retângulo: ")
    fmt.Scan(&base)

    fmt.Print("Digite a altura do retângulo: ")
    fmt.Scan(&altura)

    area := geometria.Area(base, altura)
    perimetro := geometria.Perimetro(base, altura)

    fmt.Printf("Área do retângulo: %.2f\n", area)
    fmt.Printf("Perímetro do retângulo: %.2f\n", perimetro)

}



