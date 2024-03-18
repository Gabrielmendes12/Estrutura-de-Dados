package main

import "fmt"

const M = 5

func exibir(pilha *[M]int, topo *int){
	for i:= 0; i <= *topo; i++{
		fmt.Printf("%d", pilha[i])
	}
	fmt.Println("")
}

func push(pilha *[M]int, topo *int, valor int) {
	if *topo == M - 1 {
		fmt.Println("Overflow")
	} else {
		*topo++
		pilha[*topo] = valor
	}
}

/*func pop(pilha *[M]int, topo *int, valor int) error {
	if
} */

// Nas pilhas a adição de um elemento smp aparecerá no topo desta pilha
func main(){
	var pilha [M]int
	var topo = -1 // pilha está vazia inicialmente

	exibir(&pilha, &topo)  // []
	push(&pilha, &topo, 2) // [2]
	push(&pilha, &topo, 3) // [2, 3]
	push(&pilha, &topo, 4) // [2, 3, 4]
	exibir(&pilha, &topo) // [2, 3, 4]
	//pop(pilha) // [2, 3]
	//pop(pilha) // [2, 3]
	//exibir(pilha)
}