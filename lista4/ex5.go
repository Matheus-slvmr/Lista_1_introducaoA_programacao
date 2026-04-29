/* Faça um programa que leia um vetor de inteiros, de 10 posições. A seguir, encontre o menor elemento (X) do
vetor. Imprima uma mensagem mostrando: “O menor elemento do vetor é”, X, “e sua posição dentro do vetor
é:”,P. Assuma que os elementos informados no vetor são todos diferentes entre si./*  */
package main

import "fmt"

func main() {
    var vetor [10]int
	var menor int
	var posicao int
    // Leitura
    for i := 0; i < 10; i++ {
        fmt.Scan(&vetor[i])
    }
	menor = vetor[0]
	posicao = 0

    // Processamento
    for i := 0; i < 10; i++ {
        if vetor[i] < menor {
            menor = vetor[i]
			posicao = i
        }
    }
	fmt.Printf("o menor é %d a posição é %d", menor, posicao)
}