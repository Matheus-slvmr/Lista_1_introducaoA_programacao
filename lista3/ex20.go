package main

import (
	"fmt"
)

func main() {
	fmt.Println("Índices abaixo da diagonal principal (matriz 10x10):")

	// Loop externo: controla as linhas (i)
	for i := 0; i < 10; i++ {
		
		// Loop interno: controla as colunas (j)
		// O truque aqui: j começa em 0, mas para de rodar assim que encosta no valor de i
		for j := 0; j < i; j++ {
			fmt.Printf("[%d][%d] ", i, j)
		}
		
		// Quebra de linha para visualizar o formato do triângulo inferior
		fmt.Println()
	}
}