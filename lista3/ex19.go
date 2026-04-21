package main

import (
	"fmt"
)

func main() {
	fmt.Println("Índices acima da diagonal principal (matriz 10x10):")

	// Loop externo: controla as linhas (i)
	for i := 0; i < 10; i++ {
		
		// Loop interno: controla as colunas (j)
		// O pulo do gato: j não começa em 0. Ele começa uma casa APÓS a linha atual (i + 1)
		for j := i + 1; j < 10; j++ {
			fmt.Printf("[%d][%d] ", i, j)
		}
		
		// Quebra de linha para você visualizar o formato de "triângulo" que vai se formar
		fmt.Println()
	}
}