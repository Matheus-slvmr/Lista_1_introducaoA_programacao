package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Print("Digite a quantidade de termos (N): ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Erro: O número de termos (N) deve ser maior que zero.")
		return 
	}

	var somaTotal float64
	numerador := 1000.0 

	
	for i := 1; i <= n; i++ {
		
		
		denominador := float64(i)
		
	
		termoAtual := numerador / denominador

		// Regra do Sinal: Par subtrai, Ímpar soma
		if i%2 == 0 {
			somaTotal -= termoAtual 
		} else {
			somaTotal += termoAtual 
		}

		numerador -= 3.0
	}

	
	fmt.Printf("\nO resultado da série com %d termos é: %.4f\n", n, somaTotal)
}