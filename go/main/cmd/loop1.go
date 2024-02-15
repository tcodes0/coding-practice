package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Definição da estrutura Aluno
type Aluno struct {
	nome string // Nome do aluno
	nota float64 // Nota do aluno
}

// Definição da função EscolaSP
func EscolaSP() {
	// Imprime linhas de separação
	fmt.Println("--------------------")
	fmt.Println("Escola Santa Paciencia") // Imprime o nome da escola
	fmt.Println("--------------------")
	fmt.Println("Quantos alunos a turma tem?") // Solicita o número de alunos na turma

	// Preparação para ler a entrada do usuário
	reader := bufio.NewReader(os.Stdin)
	rawAlunos, err := reader.ReadString('\n') // Lê a entrada do usuário
	if err != nil {
		log.Fatal(err)
	}
	reader.Reset(os.Stdin)

	rawAlunos = strings.Replace(rawAlunos, "\n", "", -1) // Remove a quebra de linha
	n, err := strconv.ParseInt(rawAlunos, 10, 10) // Converte a entrada do usuário para um número inteiro
	if err != nil {
		log.Fatal(err)
	}

	// Cria um slice de Aluno com base no número de alunos fornecido
	alunos := make([]Aluno, int(n))

	fmt.Println("--------------------")

	var melhorAluno *Aluno // Variável para armazenar o melhor aluno
	for i, aluno := range alunos {
		fmt.Printf("ALUNO %d\n", i+1) // Imprime o número do aluno
		fmt.Print("Nome do aluno: ")

		// Lê o nome do aluno fornecido pelo usuário
		rawNome, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		rawNome = strings.Replace(rawNome, "\n", "", -1) // Remove a quebra de linha

		aluno.nome = rawNome // Define o nome do aluno

		fmt.Printf("Nota de %s: ", aluno.nome) // Solicita a nota do aluno
		rawNota, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		rawNota = strings.Replace(rawNota, "\n", "", -1) // Remove a quebra de linha

		// Converte a nota fornecida pelo usuário para um número decimal
		nota, err := strconv.ParseFloat(rawNota, 64)
		if err != nil {
			log.Fatal(err)
		}

		aluno.nota = nota // Define a nota do aluno

		// Verifica se o aluno atual tem a melhor nota
		if melhorAluno == nil {
			melhorAluno = &aluno
		} else if aluno.nota > melhorAluno.nota {
			melhorAluno = &aluno
		}

		fmt.Println("--------------------")
		reader.Reset(os.Stdin)
	}

	// Imprime a informação do aluno com a melhor nota ou indica que ninguém tirou nota maior que zero
	if melhorAluno != nil {
		fmt.Printf("A melhor aproveitamento foi de %s com a nota %f\n", melhorAluno.nome, melhorAluno.nota)
	} else {
		fmt.Println("Escola nao tem alunos!")
	}
}

