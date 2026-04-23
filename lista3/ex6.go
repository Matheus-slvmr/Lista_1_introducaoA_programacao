/* Escreva um programa que receba um número inteiro positivo, verifique e informe se ele é ou não um número
triangular. Obs.: Um número é triangular quando é resultado do produto de três números naturais consecutivos.
Exemplo: 24 = 2 x 3 x 4; 120 = 4 x 5 x 6
*/
package main
import "fmt"
func main() {
	var numero int
	fmt.Print("Digite um número inteiro positivo: ")
	fmt.Scan(&numero)
	if numero <= 0 {
		fmt.Println("Por favor, digite um número inteiro positivo.")
		return
	}
	triangular := false
	for i := 1; i*(i+1)*(i+2) <= numero; i++ {
		if i*(i+1)*(i+2) == numero {
			triangular = true
			break
		}
	}
	if triangular {
		fmt.Println("O número é triangular.")
	} else {
		fmt.Println("O número não é triangular.")
	}
}

