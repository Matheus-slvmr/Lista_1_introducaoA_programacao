package main

import "fmt"

func main() {
	var vetorA [10]int
	contagem := make(map[int]int)

	// 1. Leitura dos 10 números
	fmt.Println("Digite 10 números inteiros:")
	for i := 0; i < 10; i++ {
		fmt.Printf("Posição %d: ", i)
		fmt.Scan(&vetorA[i])

		// 2. Preenchimento do mapa de frequências
		// Incrementa o valor associado à chave (o número digitado)
		contagem[vetorA[i]]++
	}

	// 3. Verificação e exibição dos repetidos
	encontrouRepetido := false
	fmt.Println("\n--- Elementos Repetidos ---")
	
	for numero, qtd := range contagem {
		if qtd > 1 {
			fmt.Printf("O número %d se repete %d vezes.\n", numero, qtd)
			encontrouRepetido = true
		}
	}

	if !encontrouRepetido {
		fmt.Println("Nenhum número se repetiu.")
	}
}