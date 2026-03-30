/*Um comerciante calcula o valor da venda, tendo em vista a tabela a seguir:
Valor da Compra Valor da Venda
Valor < R$ 10,00                   Lucro de 70%
R$ 10,00 <= Valor < R$ 30,00       Lucro de 50%
R$ 30,00 <= Valor < R$ 50,00 	   Lucro de 40%
Valor >= 50,00 					   Lucro de 30%

Escreva um programa que leia o valor da compra e imprima o valor da venda. Cuidado com valor
inválido de compra!
*/
package main
import (
	f "fmt"
)
func main (){
	compra , venda := 0.0 , 0.0
f.Print("digite o valor da compra: ")
f.Scan(&compra)

if compra < 10.0 {
	venda = compra + ( compra * 0.7)
	f.Printf("o valor da venda é : %v", venda)
	}else if compra >= 10.0 && compra < 30.0 {
	venda = compra + (compra * 0.5)
	f.Printf("o valor da venda é : %v", venda)
		}else if compra >= 30.0 && compra < 50.0 {
	venda = compra + (compra * 0.4)
	f.Printf("o valor da venda é : %v", venda)
			}else if compra >= 50.0 {
	venda = compra + (compra * 0.3)
	f.Printf("o valor da venda é : %v", venda)
}
}