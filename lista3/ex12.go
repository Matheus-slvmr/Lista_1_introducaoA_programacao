package main

import (
	f "fmt"
)

// Função auxiliar para calcular o fatorial de forma precisa para a série
func calcularFatorial(n int) float64 {
	fat := 1.0
	for i := 1; i <= n; i++ {
		fat *= float64(i)
	}
	return fat
}

func main() {
	//
	var x float64
	f.Print("Digite o valor de X: ")
	f.Scan(&x)

	var s float64 = x // O primeiro termo da série é o próprio X
	
	// Como já temos o primeiro termo (X), vamos calcular os próximos 19 
	// para completar os 20 termos solicitados.
	for i := 1; i < 20; i++ {
		termo := x / calcularFatorial(i)

		// Se o índice for ímpar (1, 3, 5...), subtraímos
		// Se o índice for par (2, 4, 6...), somamos
		if i%2 != 0 {
			s -= termo
		} else {
			s += termo
		}
	}

	f.Printf("\nO resultado do somatório S para os 20 primeiros termos é: %.10f\n", s)
}