package main

import "fmt"

func main() {
    var codigoPeca1, numeroPecas1, codigoPeca2, numeroPecas2 int
    var valorUnitario1, valorUnitario2 float64

    // Lendo os dados da primeira peça
    fmt.Scanln(&codigoPeca1, &numeroPecas1, &valorUnitario1)

    // Lendo os dados da segunda peça
    fmt.Scanln(&codigoPeca2, &numeroPecas2, &valorUnitario2)

    // Calculando o valor total a ser pago
    totalPagar := (float64(numeroPecas1) * valorUnitario1) + (float64(numeroPecas2) * valorUnitario2)

    // Imprimindo o resultado
    fmt.Printf("VALOR A PAGAR: R$ %.2f\n", totalPagar)
}
