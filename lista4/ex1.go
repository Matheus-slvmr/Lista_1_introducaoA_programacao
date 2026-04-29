package main

import "fmt"

func main() {
    var vetor [10]int

    // Leitura
    for i := 0; i < 10; i++ {
        fmt.Scan(&vetor[i])
    }

    // Processamento
    for i := 0; i < 10; i++ {
        if vetor[i] > 50 {
            fmt.Printf("Valor %d na posição %d\n", vetor[i], i)
        }
		if vetor[i] < 50 {
		fmt.Print("nao existe\n")
	}
    }
	
}