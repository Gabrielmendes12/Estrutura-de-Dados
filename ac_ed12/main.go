package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
)

func alarme(){
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
	
		line := scanner.Text()
		// Dividindo em componentes
		parts := strings.Fields(line)

		// Convertendo as partes em inteiros
		h1, _ := strconv.Atoi(parts[0])
		m1, _ := strconv.Atoi(parts[1])
		h2, _ := strconv.Atoi(parts[2])
		m2, _ := strconv.Atoi(parts[3])

		if h1 == 0 && m1 == 0 && h2 == 0 && m2 == 0 {
			break
		}

		// Calculando tempo em minutos
		time1 := h1*60 + m1
		time2 := h2*60 + m2

		// Calculando a diferença em minutos
		diff := time2 - time1

		// Se a diferença é negativa, isso significa que o alarme foi configurado pro dia seguinte
		if diff < 0 {
			diff += 24 * 60
		}
		fmt.Println(diff)
	}
}

func topN(){
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		// Convertendo o input em inteiro
		K, err := strconv.Atoi(line)
		if err != nil {
			continue
		}

		// Determinando a categoria
		if K <= 1 {
			fmt.Println("Top 1")
		} else if K <= 3 {
			fmt.Println("Top 3")
		} else if K <= 5 {
			fmt.Println("Top 5")
		} else if K <= 10 {
			fmt.Println("Top 10")
		} else if K <= 25 {
			fmt.Println("Top 25")
		} else if K <= 50 {
			fmt.Println("Top 50")
		} else if K <= 100 {
			fmt.Println("Top 100")
		}
	}
}


func escada() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// Lendo o número de passos
		nStr := scanner.Text()
		n, err := strconv.Atoi(nStr)
		if err != nil {
			break
		}

		if !scanner.Scan() {
			break
		}

		// Lendo as dimensões pra cada passo
		dimensions := strings.Fields(scanner.Text())
		if len(dimensions) < 3 {
			break
		}
		h, _ := strconv.Atoi(dimensions[0])
		c, _ := strconv.Atoi(dimensions[1])
		l, _ := strconv.Atoi(dimensions[2])

		// Usando a variável para evitar erro de compilação
		_ = h 

		// Cálculo da área
		areaCm2 := float64(n * c * l)
		areaM2 := areaCm2 / 10000.0


		fmt.Printf("%.4f\n", round(areaM2, 4))
	}
}

func round(val float64, precision int) float64 {
    factor := math.Pow(10, float64(precision))
    return math.Round(val*factor) / factor
}


func main() {
	//alarme()
	//topN()
	//escada()
}
