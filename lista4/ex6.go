/* Escreva um programa que armazene todos os números inteiros de 100 a 1 (ordem decrescente) em um vetor. A
seguir, imprima os elementos do vetor. */
package main

import "fmt"

func main() {
    var vetor [100]int

    // Leitura
    for i := 0; i < 100; i++ {
		vetor[i] = 100 - i
		fmt.Print(vetor[i],"\n")
    }

	
}