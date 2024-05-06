package main

import (
	"fmt"
)

// Fila: elemento inserido no final e elemento removido no início

type No struct {
	valor int
	prox *No
}

type Fila struct {
	tamanho int
	inicio  *No1
	fim     *No2
}

// Percorrendo a fila

func (f *Fila) percorrer(){
	no := f.inicio
	if no == nil {
		fmt.Println("fila vazia")
	} else {
		for no.prox != nil {
			fmt.Print(no.valor, " -> ")
			no = no.prox
		}
		fmt.Println(no.valor)
	}
}

func (f *Fila) inserirFinal(){
	


}

func main(){
	fmt.Println("hello world")
}
