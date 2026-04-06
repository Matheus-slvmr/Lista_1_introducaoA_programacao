package main
import (
	"fmt"
)
type Pessoa struct {
	nome      string
	idade     int
	profissao string
	salario   int
}
func main() {
	var pes1 Pessoa
	// Especificacao da pes1
	pes1.nome = "Monica"
	pes1.idade = 35
	pes1.profissao = "Engenheiro de Software"
	pes1.salario = 7000
	// Imprimir as informacoes da pessoa1
	fmt.Println("Informacoes da pessoa 1:")
	fmt.Println("nome: ", pes1.nome)
	fmt.Println("idade: ", pes1.idade)
	fmt.Println("profissao: ", pes1.profissao)
	fmt.Println("salario: ", pes1.salario)
	//pessoa 2
	var pes2 Pessoa
	pes2.nome = "Joao"
	pes2.idade = 28
	pes2.profissao = "Designer Grafico"
	pes2.salario = 4500
	fmt.Println("\nInformacoes da pessoa 2:")
	fmt.Println("nome: ", pes2.nome)
	fmt.Println("idade: ", pes2.idade)
	fmt.Println("profissao: ", pes2.profissao)
	fmt.Println("salario: ", pes2.salario)
}
