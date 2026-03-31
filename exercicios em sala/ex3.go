package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Print("Digite a quantidade de números inteiros que deseja inserir: ")
	fmt.Scan(&n) 

	numeros := make([]int, n)
	fmt.Println("Digite", n, "números inteiros:")

	for i := 0; i < n; i++ {
		fmt.Printf("%dº número: ", i+1)
		fmt.Scan(&numeros[i])
	}

	fmt.Println("\nExibindo na ordem inversa:")

	//lista de tras pra frente
	for i := n - 1; i >= 0; i-- {
		fmt.Printf("%d ", numeros[i])
	}
	
}