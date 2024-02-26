package main

import "fmt"

// Exercício 1

type Contato struct {
	Nome string
	Email string
}


// Método para alterar o e-mail do Contato
func (c *Contato) AlterarEmail(novoEmail string) {
    c.Email = novoEmail
}

// Função para adicionar um Contato ao array
func adicionarContato(contato Contato, arrayContatos []Contato) []Contato {
    for i := range arrayContatos {
        // Verifica se o elemento no índice i está vazio
        if arrayContatos[i].Nome == "" && arrayContatos[i].Email == "" {
            // Adiciona o contato no primeiro elemento vazio
            arrayContatos[i] = contato
            return arrayContatos
        }
    }
    // Se não encontrar nenhum elemento vazio, retorna o array sem modificação
    return arrayContatos
}

// Função para excluir o último contato válido do array
func excluirContato(arrayContatos []Contato) []Contato {
    // Itera reversamente pelo array para encontrar o último contato válido
    for i := len(arrayContatos) - 1; i >= 0; i-- {
        if arrayContatos[i].Nome != "" || arrayContatos[i].Email != "" {
            // Elimina o último contato válido encontrando
            arrayContatos[i] = Contato{}
            break
        }
    }
    return arrayContatos
}



func main (){

	// chamada da função do exercício 1

	// Criando uma instância do struct Contato
    contato := Contato{
        Nome:  "Gabriel",
        Email: "gabriel@gmail.com",
    }

    // Exibindo o contato antes da alteração
    fmt.Println("Contato antes da alteração:")
    fmt.Printf("Nome: %s\nEmail: %s\n", contato.Nome, contato.Email)

    // Alterando o e-mail usando o método
    contato.AlterarEmail("gabriel_novo@gmail.com")

    // Exibindo o contato após a alteração
    fmt.Println("\nContato após a alteração:")
    fmt.Printf("Nome: %s\nEmail: %s\n", contato.Nome, contato.Email)


	// chamada da função do exercício 2

	// Criando um array de Contatos com capacidade para até 5 elementos
    arrayContatos := make([]Contato, 3)

    // Criando um novo Contato
    novoContato := Contato{
        Nome:  "Fernando",
        Email: "fernando@gmail.com",
    }

	novoContato2 := Contato{
        Nome:  "Bernardo",
        Email: "bernardo@gmail.com",
    }

    novoContato3 := Contato{
        Nome:  "Alice",
        Email: "alice@gmail.com",
    }


    // Adicionando o Contato ao array
    arrayContatos = adicionarContato(novoContato, arrayContatos)
	arrayContatos = adicionarContato(novoContato2, arrayContatos)
	arrayContatos = adicionarContato(novoContato3, arrayContatos)


    // Exibindo o array de Contatos após a adição
    fmt.Println("\nArray de Contatos após a adição:")
    for _, contato := range arrayContatos {
        fmt.Printf("Nome: %s, Email: %s\n", contato.Nome, contato.Email)
	}

	// Chamada da função para excluir o último contato
	arrayContatos = excluirContato(arrayContatos)

	// Exibindo o array de Contatos após a exclusão
	fmt.Println("\nArray de Contatos após a exclusão:")
	for _, contato := range arrayContatos {
		fmt.Printf("Nome: %s, Email: %s\n", contato.Nome, contato.Email)
	}
}