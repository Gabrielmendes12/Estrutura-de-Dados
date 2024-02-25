package main

import "fmt"

// Exercício 1

func calculaMedia(valores ...float64) float64 {
	total := 0.0 // acumulará a soma dos valores

	for _, valor := range valores { // iteração sobre os valores(parâmetros) da função; for ignorando o índice
		total += valor
	}

	if len(valores) > 0 { // condição para possibilitar o cálculo da média
        return total / float64(len(valores)) // média aritmética = soma dos elementos / quantidade de elementos 
    } else {
        return 0.0
    }
}

// Exercício 2

func verificaParidade(numero int) string {
	if numero % 2 == 0 {  // se a divisão por 2 for igual a 0
		return "Par"
	} else {
		return "Ímpar"
	}
}

// Exercício 3

func minhaPotencia(base, expoente int) int {
	resultado := 1  // inicialização do cálculo = 1 pois todo número elevado a 0 é igual a 1

	for i := 0; i < expoente; i++ {
		resultado *= base // resultado é multiplicado pela base em cada iteração
	}

	return resultado
	// loop opera a multiplicação da base pelo n° de vzs do expoente calculando o resultado
}

// Exercício 4

func converteCelsiusParaFahrenheit(graus_C float64) float64 {
    // A fórmula de conversão de Graus Celsius para Graus Fahrenheit é: F = (C * 9/5) + 32
    graus_F := (graus_C * 9/5) + 32
    return graus_F
}


func main(){
	// chamada da função do exercício 1
	media := calculaMedia(4.0, 5.0, 6.5, 8.5)
	fmt.Println("A média aritmética é:", media)

	// chamada da função do exercício 2
	numero := 9
	resposta := verificaParidade(numero)
	fmt.Println("O número recebido é:",(resposta))

	// chamada da função do exercício 3
	base := 4
	expoente := 3
	resultado := minhaPotencia(base, expoente)
	fmt.Println((base),"elevado a", (expoente), "é igual a",(resultado))

	// chamada da função do exercício 4
	temperatura_C := 20.0
	temperatura_F := converteCelsiusParaFahrenheit(temperatura_C)

	fmt.Println((temperatura_C),"graus Celsius é igual a",(temperatura_F), "graus Fahrenheit")
}


