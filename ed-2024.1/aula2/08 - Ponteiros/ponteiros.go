package main

import (
	"fmt"
	"strings"
)

func main(){
x := 2
y := x

fmt.Println(x, y)

x = 3
fmt.Println(x, y)

z := &x // referência  // atribuição de ponteiro // z vai imprimir o endereço da memória de x
fmt.Println(x, z)
fmt.Println(x, *z) // deferência  // z vai acessar o valor armazenado no endereço de memória

msg1 := "olá, mundo"
alteraMensagem(&msg1)
fmt.Println(msg1)

n1 := Numero{Valor:10}
fmt.Println(n1) //10
n1.Incremento()
fmt.Println(n1) // 11


}

/*
Usamos ponteiros como parâmetros de funções quando:
1- é necessário reduzir o espaço consumido em memória
2 - queremos alterar o valor dos parâmetros
*/


func alteraMensagem(msg *string) {
	novaMensagem := strings.ReplaceAll(*msg, "mundo", "turma")
	*msg = novaMensagem  // acessa o valor de msg 1 (olá, mundo) da func main e substitui mundo por turma -> "olá, turma"
}


type Numero struct {
	Valor int
}

func (n *Numero)Incremento(){ // este ponteiro altera o valor usando o incremento
	n.Valor++
}

