package main

import "fmt"

// Função auxiliar para verificar se um número é primo
func ehPrimo(numero int) bool {
	// 0 e 1 não são números primos
	if numero <= 1 {
		return false
	}
	
	// Tentamos dividir o número por todos os valores entre 2 e ele mesmo - 1.
	// Se o resto da divisão (%) for zero em algum momento, ele NÃO é primo.
	for i := 2; i < numero; i++ {
		if numero%i == 0 {
			return false
		}
	}
	
	return true
}

func main() {
	var n1, n2 int

	fmt.Print("Digite dois numeros (separados por espaço ou enter): ")
	
	// CORREÇÃO 1: Usar o "&" (ponteiro) para salvar os valores digitados
	fmt.Scan(&n1, &n2)

	fmt.Printf("\nOs números primos entre %d e %d são:\n", n1, n2)

	for i := n1; i <= n2; i++ {
		
		// CORREÇÃO 3: Usamos a nossa função para testar o número atual (i)
		if ehPrimo(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println() // Apenas para quebrar a linha no final do terminal
}