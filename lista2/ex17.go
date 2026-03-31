/*
Desenvolver um programa para calcular a conta de água para a SANEAGO. O custo da água varia dependendo
do tipo do consumidor - residencial, comercial ou industrial. A regra para calcular a conta é:
• Residencial: R$ 5,00 de taxa mais R$ 0,05 por m3 gastos;
• Comercial: R$ 500,00 para os primeiros 80 m3 gastos mais R$ 0,25 por m3 gastos acima dos 80 m3;
• Industrial: R$ 800,00 para os primeiros 100 m3 gastos mas R$ 0,04 por m3 gastos acima dos 100 m3;
O programa deverá ler a conta do cliente, seu tipo (residencial, comercial e industrial) e o seu consumo de água em
metros cúbicos. Como resultado imprimir a conta do cliente e o valor em real a ser pago pelo mesmo.
*/
package main
import (
	"fmt"
)
func main() {
	var tipo string
	var consumo float64
	fmt.Print("Digite o tipo do consumidor (residencial, comercial ou industrial): ")
	fmt.Scan(&tipo)
	fmt.Print("Digite o consumo de água em metros cúbicos: ")
	fmt.Scan(&consumo)
	var conta float64
	switch tipo {
	case "residencial":
		conta = 5 + 0.05*consumo
	case "comercial":
		if consumo <= 80 {
			conta = 500
		} else {
			conta = 500 + 0.25*(consumo-80)
		}
	case "industrial":
		if consumo <= 100 {
			conta = 800
		} else {
			conta = 800 + 0.04*(consumo-100)
		}
	default:
		fmt.Println("Tipo de consumidor inválido.")
		return
	}
	fmt.Printf("Conta do cliente: R$ %.2f\n", conta)
}