
package main
import (
	"fmt"
)
type Pessoa struct {
	nome      string
	altura    float64
	peso 	float64
}
func main() {
	var pessoas []Pessoa
	for {
		var p Pessoa
		fmt.Print("Digite o nome da pessoa (ou 'FIM' para encerrar): ")
		fmt.Scanln(&p.nome)
		if p.nome == "FIM" {
			break
		}
		fmt.Print("Digite a altura da pessoa: ")
		fmt.Scanln(&p.altura)
		p.peso = 72.7*p.altura - 58.0
		pessoas = append(pessoas, p)
	}
	fmt.Println("Dados das pessoas:")
	for _, pessoa := range pessoas {
		fmt.Printf("Nome: %s, Altura: %.2f, Peso Ideal: %.2f\n", pessoa.nome, pessoa.altura, pessoa.peso)
	}
}
