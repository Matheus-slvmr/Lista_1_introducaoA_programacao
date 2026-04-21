package main

import (
	"fmt"
)

func main() {
	fmt.Println("Índices da diagonal principal de uma matriz 10x10:")

	// Usamos apenas uma variável (i) que servirá tanto para a linha quanto para a coluna
	for i := 0; i < 10; i++ {
		// Imprimimos o mesmo 'i' duas vezes
		fmt.Printf("[%d][%d] ", i, i)
	}
	
	fmt.Println() // Quebra de linha final
}