package main

import "fmt"

func MergeListas(a, b []int) []int {
    m, n := len(a), len(b)
    v := make([]int, m+n)
    i, j, k := 0, 0, 0

    for i < m && j < n {
        if a[i] < b[j] {
            v[k] = a[i]
            i++
        } else {
            v[k] = b[j]
            j++
        }
        k++
    }

    for i < m {
        v[k] = a[i]
        i++
        k++
    }

    for j < n {
        v[k] = b[j]
        j++
        k++
    }

    return v
}


func main(){

	 // Exemplo de listas ordenadas
	 listaA := []int{1, 2, 3, 4, 5}
	 listaB := []int{6, 7, 8, 9, 10}

	 // Chama a função MergeListas
	 listaResultado := MergeListas(listaA, listaB)

	 // Exibe o resultado
	 fmt.Println("Lista Resultado:", listaResultado)
}