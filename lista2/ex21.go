/* Escrever um programa que leia o número de identificação, as 3 notas obtidas por um aluno nas 3 verificações e
a média dos exercícios que fazem parte da avaliação. Calcular a média de aproveitamento do aluno, usando a
fórmula:
Média Final = ( nota1 + nota2 ∗ 2 + nota3 ∗ 3 + médiados exercícios )
7
e o seu conceito, usando a tabela a seguir:
Média de Aproveitamento Conceito
>= 9,0 e <= 10,0 A
>=7,5 e < 9,0 B
>=6,0 e < 7,5 C
>=4,0 e < 6,0 D
< 4,0 E
O programa deve escrever o número do aluno, suas notas, a média dos exercícios, a média de aproveitamento,
o conceito correspondente e a mensagem: APROVADO se o conceito for A, B ou C e REPROVADO, se o
conceito for D ou E.
22. Desenvolver um programa */
package main
import (
	"fmt"
)
func main() {
	var id int
	fmt.Print("Digite o número de identificação do aluno: ")
	fmt.Scan(&id)
	var nota1, nota2, nota3, media_exercicios float64
	fmt.Print("Digite a nota da primeira verificação: ")
	fmt.Scan(&nota1)
	fmt.Print("Digite a nota da segunda verificação: ")
	fmt.Scan(&nota2)
	fmt.Print("Digite a nota da terceira verificação: ")
	fmt.Scan(&nota3)
	fmt.Print("media dos exercicios: ")
	media_exercicios = nota1 + nota2 + nota3 / 3
	fmt.Printf("Média dos exercícios: %.2f\n", media_exercicios)
	media_aproveitamento := (nota1 + nota2*2 + nota3*3 + media_exercicios) / 7
	var conceito string
	if media_aproveitamento >= 9.0 && media_aproveitamento <= 10.0 {
		conceito = "A"
	} else if media_aproveitamento >= 7.5 && media_aproveitamento < 9.0 {
		conceito = "B"
	} else if media_aproveitamento >= 6.0 && media_aproveitamento < 7.5 {
		conceito = "C"
	} else if media_aproveitamento >= 4.0 && media_aproveitamento < 6.0 {
		conceito = "D"
	} else {
		conceito = "E"
	}
	fmt.Printf("Número do aluno: %d\n", id)
	fmt.Printf("Média de aproveitamento: %.2f\n", media_aproveitamento)
	fmt.Printf("Conceito: %s\n", conceito)

	if conceito == "A" || conceito == "B" || conceito == "C" {
		fmt.Println("APROVADO")
	} else {
		fmt.Println("REPROVADO")
	}
}