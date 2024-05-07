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
	inicio  *No
	fim     *No
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

// Inserindo um elemento no final da fila

func (f *Fila) inserirFinal(novoValor int){
	novoNo := &No{valor: novoValor}

	if f.inicio == nil {
		f.inicio = novoNo
		f.fim = f.inicio
	} else {
		f.fim.prox = novoNo
		f.fim = novoNo
	}
	f.tamanho++
}

// Fazendo a remoção e o retorno do elemento inicial da fila

func (f *Fila) removerInicio() (int, error) {
	if f.inicio == nil {
		return -1, fmt.Errorf("A fila está vazia")
	}

	valor := f.inicio.valor
	f.inicio = f.inicio.prox

	if f.inicio == nil {
		f.fim = nil
	}
	f.tamanho--
	return valor, nil
}

func main(){
	
	fila := Fila{}

	// Fazendo a inserção de elementos
	fila.inserirFinal(1)
	fila.inserirFinal(2)
	fila.inserirFinal(3)
	fila.inserirFinal(4)

	fmt.Println("Elementos da fila: ")
	fila.percorrer()
	fmt.Println("Tamanho da fila: ", fila.tamanho)

	// Fazendo a remoção de elementos
	valorRemovido, _ := fila.removerInicio()
	fmt.Println("Elemento removido:", valorRemovido)

	fila.percorrer()
	fmt.Println("Tamanho da fila: ", fila.tamanho)
	
	fila.inserirFinal(5)
	fmt.Println("Elementos da fila: ")
	fila.percorrer()

}
