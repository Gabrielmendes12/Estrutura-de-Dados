package main

import (
	"fmt"
)

func countDiamonds(input string) int {
	stack := make([]rune, 0) // criação da pilha

	diamonds := 0
	for _, char := range input {
		if char == '<' {
			stack = append(stack, char)
		} else if char == '>' && len(stack) > 0 {
			stack = stack[:len(stack)-1]
			diamonds++
		}
	}

	return diamonds
}

func main() {
	var N int
	fmt.Scan(&N) // casos de teste

	for i := 0; i < N; i++ {
		var input string
		fmt.Scan(&input)
		fmt.Println(countDiamonds(input))
	}
}
