/* Escreva um programa que receba a idade, a altura e o peso de várias pessoas. Calcule e imprima:
- a quantidade de pessoas com idade superior a 50 anos;
- a média das alturas das pessoas com idade entre 10 e 20 anos;
- a porcentagem de pessoas com peso inferior a 40 quilos entre todas as pessoas analisadas.
Considere que os dados informados são válidos. Pergunte ao usuário se ele deseja continuar digitando dados ou
não (Exemplo: 1 - Sim, Outro valor diferente de 1 - Não). */
package main
import f "fmt"
func main(){
var idade int
var altura float64
var peso float64
// Variável para controlar a continuação do programa
var continuar int
//
var countIdadeSuperior50 int
var countIdadeEntre10e20 int
var somaAlturaIdadeEntre10e20 float64
var countPesoInferior40 int
var countTotalPessoas int
for {
	f.Print("Digite a idade: ")
	f.Scan(&idade)
	f.Print("Digite a altura: ")
	f.Scan(&altura)
	f.Print("Digite o peso: ")
	f.Scan(&peso)

	countTotalPessoas++
	if idade > 50 {
		countIdadeSuperior50++
	}
	if idade >= 10 && idade <= 20 {
		countIdadeEntre10e20++
		somaAlturaIdadeEntre10e20 += altura
	}
	if peso < 40 {
		countPesoInferior40++
	}
	f.Print("Deseja continuar? (1 - Sim, outro valor - Não): ")
	f.Scan(&continuar)
	if continuar != 1 {
		break
	}
}
if countIdadeEntre10e20 > 0 {
	f.Printf("Média das alturas das pessoas com idade entre 10 e 20 anos: %.2f\n", somaAlturaIdadeEntre10e20/float64(countIdadeEntre10e20))
}
f.Printf("Quantidade de pessoas com idade superior a 50 anos: %d\n", countIdadeSuperior50)
if countTotalPessoas > 0 {
	f.Printf("Porcentagem de pessoas com peso inferior a 40 quilos: %.2f%%\n", float64(countPesoInferior40)/float64(countTotalPessoas)*100)
}
}