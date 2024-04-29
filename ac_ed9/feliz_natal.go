package main

import (
	"fmt"
	"strings"
)

// Função para obter a frase de Natal no idioma correto dado o nome do país
func obterFraseNatal(pais string) string {
	// Mapa de frases de Natal por país
	frasesNatal := map[string]string{
		"brasil":         "Feliz Natal!",
		"alemanha":       "Frohliche Weihnachten!",
		"austria":        "Frohe Weihnacht!",
		"coreia":         "Chuk Sung Tan!",
		"espanha":        "Feliz Navidad!",
		"grecia":         "Kala Christougena!",
		"estados-unidos": "Merry Christmas!",
		"inglaterra":     "Merry Christmas!",
		"australia":      "Merry Christmas!",
		"portugal":       "Feliz Natal!",
		"suecia":         "God Jul!",
		"turquia":        "Mutlu Noeller!",
		"argentina":      "Feliz Navidad!",
		"chile":          "Feliz Navidad!",
		"mexico":         "Feliz Navidad!",
		"antardida":      "Merry Christmas!",
		"canada":         "Merry Christmas!",
		"irlanda":        "Nollaig Shona Dhuit!",
		"belgica":        "Zalig Kerstfeest!",
		"italia":         "Buon Natale!",
		"libia":          "Buon Natale!",
		"siria":          "Milad Mubarak!",
		"marrocos":       "Milad Mubarak!",
		"japao":          "Merii Kurisumasu!",
	}

	// Converte o nome do país para minúsculas
	pais = strings.ToLower(pais)

	// Verifica se o país está no mapa de frases
	frase, encontrado := frasesNatal[pais]
	if encontrado {
		return frase
	}
	// Se não estiver no mapa, retorna a mensagem de not found
	return "--- NOT FOUND ---"
}

func main() {
	// Lista de países simulando os dados informados pelo painel de navegação do trenó
	paises := []string{"uri-online-judge", "alemanha", "brasil", "austria"}

	// Itera sobre a lista de países e exibe a frase de Natal correspondente
	for _, pais := range paises {
		frase := obterFraseNatal(pais)
		fmt.Println(frase)
	}
}