package main

import (
	"fmt"
)

func main() {
	var somaTotal float64

	fmt.Println("Calculando a série S...")

	for i := 1.0; i <= 37.0; i++ {
		
		fator1 := 38.0 - i
		fator2 := 39.0 - i
		

		divisor := i

		
		termoAtual := (fator1 * fator2) / divisor

		
		somaTotal += termoAtual
	}

	fmt.Printf("\nO valor total da soma (S) é: %.2f\n", somaTotal)
}