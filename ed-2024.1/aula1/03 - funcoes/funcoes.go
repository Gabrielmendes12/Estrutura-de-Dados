package main

import "fmt"

func soma (a, b int) int {
	return a + b
}

func informalidade (nome string, idade int) {
	fmt.Println (nome, "tem", idade, "anos")
}

func trocaValores (x, y float64) (float64, float64) {
	return y, x
}

func main(){
	fmt.Println(soma(4,8)) //12
	informalidade("Alice", 25) // Alice tem 25 anos

	a, b := 4.4 , 8.5
	fmt.Println(trocaValores(a,b)) // 8.5 , 4.4

	anonima()

}

func anonima() {
	dobra := func(x int) (int){
		return x * 2
	}

	fmt.Println(dobra(5)) //10

	calcular := func( f func(int) int, x int) int {
		return f(x)
	}

	fmt.Println(calcular(dobra, 8))
}



