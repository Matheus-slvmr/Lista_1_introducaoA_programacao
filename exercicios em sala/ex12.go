package main

import (
	"fmt"
	"sort" // Necessário para ordenar a lista
)

func buscaBinaria(lista []int, elemento int) int {
	e := 0
	d := len(lista) - 1 // Alterado de 'n' para 'len(lista)'

	for e <= d {
		m := (e + d) / 2
		
		if lista[m] == elemento {
			return m
		} else if lista[m] < elemento {
			e = m + 1
		} else {
			d = m - 1
		}
	}
	return -1
}

func main() {
	var n, elemento int
	fmt.Print("Digite o número de elementos da lista: ")
	fmt.Scan(&n)

	lista := make([]int, n)
	fmt.Println("Digite os elementos da lista (em qualquer ordem):")
	for i := 0; i < n; i++ {
		fmt.Scan(&lista[i])
	}

	fmt.Print("Digite o elemento a ser buscado: ")
	fmt.Scan(&elemento)

	// A busca binária EXIGE que a lista esteja ordenada
	sort.Ints(lista)
	
	fmt.Println("Lista ordenada para busca:", lista)

	indice := buscaBinaria(lista, elemento)
	
	if indice != -1 {
		fmt.Printf("Elemento %d encontrado no índice: %d\n", elemento, indice)
	} else {
		fmt.Printf("Elemento %d não encontrado na lista.\n", elemento)
	}
} 