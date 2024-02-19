package main

import "fmt"

func main (){
	var x = 2
	y := 4

	// operadores aritméticos
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y) // imprime o resto da divisão inteira -> 2

	// operadores de comparação
	fmt.Println(x > y)
	fmt.Println(x < y)
	fmt.Println(x >= y)
	fmt.Println(x <= y)
	fmt.Println(x == y)
	fmt.Println(x != y)

	var a,b = true, false
	// operadores lógicos
	fmt.Println(a && b) // operador AND
	fmt.Println(a || b) // operador OR
	fmt.Println(!b) // operador NOT
}