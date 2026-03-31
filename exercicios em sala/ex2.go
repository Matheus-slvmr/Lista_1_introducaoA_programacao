//Faça um programa que leia 5 valores inteiros, armazeno-os em um vetor, calcule e apresente a soma destes valores. 
package main
import (
	"fmt"
)
func main() {
	const numeros int = 5 
	var(
		nota[numeros]float64
		soma float64
	)
	for i := 0; i < numeros; i++ {
		fmt.Printf("Digite o número %d: ", i)
		fmt.Scan(&nota[i])
		soma += nota[i]
	}
	fmt.Printf("Soma dos números: %.2f\n", soma)
}