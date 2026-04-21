package main

import (
	"fmt"
)

func main() {
	var b, n int

	fmt.Println("--- Calculadora de Exponenciação (b^n) ---")

	
	fmt.Print("Digite o valor da base (b >= 2): ")
	fmt.Scan(&b)

	fmt.Print("Digite o valor do expoente (n > 1): ")
	fmt.Scan(&n)

	if b < 2 {
		fmt.Println("Erro: A base (b) deve ser maior ou igual a 2.")
		return
	}
	if n <= 1 {
		fmt.Println("Erro: O expoente (n) deve ser maior que 1.")
		return
	}

	
	resultado := 1

	
	for i := 1; i <= n; i++ {
		resultado *= b 

	fmt.Printf("\nO valor de %d elevado a %d é: %d\n", b, n, resultado)
}