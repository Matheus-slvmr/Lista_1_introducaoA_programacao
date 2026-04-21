package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Digite a quantidade de termos (N) que deseja gerar: ")
	
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Por favor, digite um número inteiro positivo maior que zero.")
		return
	}

	fmt.Printf("Os primeiros %d termos da série são:\n", n)

	for i := 1; i <= n; i++ {
		// Calcula o quadrado da posição atual
		termo := i * i
		
	
		if i == n {
			fmt.Printf("%d\n", termo)
		} else {
			fmt.Printf("%d, ", termo)
		}
	}
}