package main

import (
	"fmt"
)

// Exercício 2

func Triangulo(){
	var A, B, C float64
    
    fmt.Scan(&A, &B, &C) // input 
	fmt.Scanln() // Limpeza do buffer de entrada

	// Lado A deve ser o maior dos 3 lados -> Os lados serão postos em ordem decrescente (A > B > C)

	// Atribuindo A = 7.0 , B = 5.0 , C = 7.0  -> Ao final devemos ter (A = 7.0 > B = 7.0 > C = 5.0)
    
    if A < B {
        A, B = B, A  // troca de valores para ordenar os lados do triângulo
    }
    if A < C {
        A, C = C, A 
    }
    if B < C {
        B, C = C, B
    }
    
    if A >= B + C {
        fmt.Println("NAO FORMA TRIANGULO")
        return // interrompe a execução para q n seja necessário verificar as condições abaixo

    } else if A*A == B*B + C*C {
        fmt.Println("TRIANGULO RETANGULO")
	
    } else if A*A > B*B + C*C {
        fmt.Println("TRIANGULO OBTUSANGULO")
	
    } else if A*A < B*B + C*C {
        fmt.Println("TRIANGULO ACUTANGULO") 
    }
    
    if A == B && B == C {
        fmt.Println("TRIANGULO EQUILATERO")

    } else if (A == B && B != C) || (A == C && C != B) || (B == C && C != A) {
        fmt.Println("TRIANGULO ISOSCELES")
    } 

}


// Exercício 3

// 2 cidades A e B -> A é sempre a cidade menor que deve ultrapassar a cidade B em população (Quantos anos levará?)

func crescimentoPopulacional() {
	var T int
	fmt.Scanln(&T)

	for i := 0; i < T; i++ {
		var PA, PB int
		var G1, G2 float64
		fmt.Scanln(&PA, &PB, &G1, &G2)

		anos := 0
		for PA <= PB {
			PA += int(float64(PA) * (G1 / 100))
			PB += int(float64(PB) * (G2 / 100))
			anos++
			if anos > 100 {
				fmt.Println("Mais de 1 seculo.")
				break
			}
		}
		if anos <= 100 {
			fmt.Println(anos, "anos.")
		}
	}
}


func main(){
	
	// Exercício 1
	fmt.Println("Hello World!")

	// chamada da função do exercício 2:
	fmt.Println()
	Triangulo()
	
	// chamada da função do exercício 3:
	fmt.Println()
	crescimentoPopulacional()
}

	

