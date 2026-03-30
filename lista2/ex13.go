/*
Escreva um programa que receba um número inteiro positivo de 3 casas e imprima o algarismo da casa das
dezenas. Não se esqueça de testar para ver se o número informado tem realmente 3 casas.
*/
package main
import f "fmt"
func main (){

	num := 0
	f.Scan(&num)
	if !(num > 99 && num < 999) {
		f.Print("numeor invalido")
	}else {
		dezena := (num / 10) % 10
		f.Printf("O algarismo da casa das dezenas de %v é: %v", num, dezena)
	}


}