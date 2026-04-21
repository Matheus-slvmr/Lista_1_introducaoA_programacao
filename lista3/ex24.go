package main

import (
	"fmt"
)

func main() {
	var inicio, fim, passo float64

	fmt.Println("--- Tabela de Seno (Série de Maclaurin Truncada) ---")
	
	// Coletando os dados do usuário
	fmt.Print("Digite o ângulo inicial (em radianos): ")
	fmt.Scan(&inicio)

	fmt.Print("Digite o ângulo final (em radianos): ")
	fmt.Scan(&fim)

	fmt.Print("Digite o passo de incremento (ex: 0.1): ")
	fmt.Scan(&passo)

	// Proteção contra loops infinitos
	if passo <= 0 {
		fmt.Println("Erro: O passo deve ser maior que zero.")
		return
	}

	// Imprimindo o cabeçalho da tabela
	fmt.Println("\n-----------------------------------------")
	fmt.Printf("%-15s | %-15s\n", "Ângulo (rad)", "Seno Aprox.")
	fmt.Println("-----------------------------------------")

	// Loop para gerar cada linha da tabela
	for A := inicio; A <= fim; A += passo {
		
		// OTIMIZAÇÃO: Em vez de usar math.Pow, calculamos as potências 
		// multiplicando a base passo a passo. É mais rápido!
		A2 := A * A
		A3 := A2 * A
		A5 := A3 * A2
		A7 := A5 * A2

		// Aplicando a fórmula exata da imagem
		senA := A - (A3 / 6.0) + (A5 / 120.0) - (A7 / 5040.0)

		// Imprimindo os valores formatados e alinhados
		fmt.Printf("%-15.4f | %-15.6f\n", A, senA)
	}
	
	fmt.Println("-----------------------------------------")
}