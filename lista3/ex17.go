package main

import (
	"fmt"
)

func main() {
	fmt.Println("Índices de uma matriz 10x10:")


	for i := 0; i < 10; i++ {
		

		for j := 0; j < 10; j++ {
			// Imprime o índice no formato [linha][coluna]
			// Usamos Printf sem quebra de linha (\n) para manter na mesma linha
			fmt.Printf("[%d][%d] ", i, j)
		}
		
		// Após terminar todas as colunas de uma linha, quebramos a linha
		// para começar a próxima linha da matriz embaixo
		fmt.Println()
	}
}