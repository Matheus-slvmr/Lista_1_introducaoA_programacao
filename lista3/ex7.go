
package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64
	var soma, somaPares float64
	var qtdTotal, qtdPares, qtdImpares int
	maior := math.Inf(-1) // Inicializa com o menor valor possível
	menor := math.Inf(1)  // Inicializa com o maior valor possível

	const FLAG = 30000

	fmt.Println("Digite os números ou (digite 30000 para finalizar):")

	for {
		fmt.Print("Número: ")
		fmt.Scan(&num)

		// Condição de parada
		if num == FLAG {
			break
		}

		// Cálculos básicos
		soma += num
		qtdTotal++

		// Maior e Menor
		if num > maior {
			maior = num
		}
		if num < menor {
			menor = num
		}

		// Lógica para Pares e Ímpares
		// Convertendo para int para usar o operador de resto %
		if int(num)%2 == 0 {
			somaPares += num
			qtdPares++
		} else {
			qtdImpares++
		}
	}

	// Verificação para evitar divisão por zero caso o usuário pare de imediato
	if qtdTotal > 0 {
		mediaGeral := soma / float64(qtdTotal)
		
		var mediaPares float64
		if qtdPares > 0 {
			mediaPares = somaPares / float64(qtdPares)
		}

		percentualImpares := (float64(qtdImpares) / float64(qtdTotal)) * 100

		fmt.Printf("a) Soma dos números: %.2f\n", soma)
		fmt.Printf("b) Quantidade de números: %d\n", qtdTotal)
		fmt.Printf("c) Média dos números: %.2f\n", mediaGeral)
		fmt.Printf("d) Maior número: %.2f\n", maior)
		fmt.Printf("e) Menor número: %.2f\n", menor)
		fmt.Printf("f) Média dos números pares: %.2f\n", mediaPares)
		fmt.Printf("g) Percentagem de ímpares: %.2f%%\n", percentualImpares)
	} else {
		fmt.Println("Nenhum número foi processado.")
	}
}