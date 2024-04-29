package main
 
import (
    "fmt"
)
 
func decryptCesarCipher(encoded string, shift int) string {
	decoded := ""
	for _, char := range encoded {
		// Convertendo o caractere para sua posição relativa em relação ao 'A'
		relativePos := int(char - 'A')
		// Aplicando o deslocamento para a esquerda (subtraindo)
		originalPos := (relativePos - shift + 26) % 26
		// Convertendo a posição de volta para o caractere
		originalChar := 'A' + rune(originalPos)
		decoded += string(originalChar)
	}
	return decoded
}

func main() {
	var N int
	fmt.Scanln(&N)

	for i := 0; i < N; i++ {
		var encoded string
		var shift int
		fmt.Scanln(&encoded)
		fmt.Scanln(&shift)

		decoded := decryptCesarCipher(encoded, shift)
		fmt.Println(decoded)
	}
}