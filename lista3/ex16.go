package main

import (
	"fmt"
)

func main() {
	var a1, a2, n int

	// Coletando os dados do usuário
	fmt.Print("Digite o 1º termo (A1): ")
	fmt.Scan(&a1)

	fmt.Print("Digite o 2º termo (A2): ")
	fmt.Scan(&a2)

	fmt.Print("Digite a quantidade de termos a gerar (N >= 3): ")
	fmt.Scan(&n)

	// Validação de segurança, caso o usuário digite um N inválido
	if n < 3 {
		fmt.Println("Por favor, digite um N maior ou igual a 3.")
		return
	}

	fmt.Printf("\nSérie de Fetuccine com %d termos:\n", n)
	
	// Imprimimos os dois primeiros termos que já conhecemos
	fmt.Printf("%d, %d", a1, a2)

	// O loop começa do 3º termo e vai até o N-ésimo
	for i := 3; i <= n; i++ {
		var proximo int

		// Se a posição (i) for par: Ai = Ai-1 - Ai-2
		if i%2 == 0 {
			// a2 é o termo anterior (Ai-1) e a1 é o retrasado (Ai-2)
			proximo = a2 - a1
		} else { 
			// Se a posição (i) for ímpar: Ai = Ai-1 + Ai-2
			proximo = a2 + a1
		}

		fmt.Printf(", %d", proximo)

		// O pulo do gato: preparamos as variáveis para a próxima rodada
		// O que era o último termo (a2) vira o penúltimo (a1)
		// O que acabamos de calcular (proximo) vira o último (a2)
		a1 = a2
		a2 = proximo
	}
	
	fmt.Println() // Apenas para quebrar a linha no final
}