/* Elabore um programa que calcule o valor a ser pago por um produto considerando o preço normal de etiqueta
e a escolha da condição de pagamento. Utilize os códigos da tabela a seguir para saber qual a condição de
pagamento escolhida e efetuar o cálculo adequado.
Código                 Condição de Pagamento
1                  	    À vista, dinheiro ou cheque, 10% de desconto
2 						À vista, cartão de crédito, 5% de desconto
3 						Em 2 vezes, preço normal de etiqueta sem juros
4 						Em 3 vezes, preço normal de etiqueta + 10% de juros 
*/
package main
import (
	"fmt"
)
func main() {	
	var preco_normal float64
	fmt.Print("Digite o preço normal do produto: ")
	fmt.Scan(&preco_normal)
	fmt.Print("Digite a condição de pagamento (1, 2, 3 ou 4): ")
	fmt.Print("1 - À vista, dinheiro ou cheque (10% de desconto)\n2 - À vista, cartão de crédito (5% de desconto)\n3 - Em 2 vezes (preço normal de etiqueta sem juros)\n4 - Em 3 vezes (preço normal de etiqueta + 10% de juros)\n")
	var condicao int
	fmt.Scan(&condicao)
	var preco_final float64
	switch condicao {
		case 1:
			preco_final = preco_normal * 0.9 
		case 2:
			preco_final = preco_normal * 0.95
		case 3:
			preco_final = preco_normal
		case 4:
			preco_final = preco_normal * 1.1
		default:
			fmt.Println("Condição de pagamento inválida.")
			return
	}
	fmt.Printf("O preço final do produto é: %.2f\n", preco_final)
}