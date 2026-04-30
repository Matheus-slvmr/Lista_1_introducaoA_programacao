package main

import "fmt"

func main() {
    var vetor1 [10]int
    var vetor2 [10]int
	var resultante [20]int

    // Leitura
    for i := 0; i < 10; i++ {
        fmt.Scan(&vetor1[i])
    }
	 for i := 0; i < 10; i++ {
        fmt.Scan(&vetor2[i])
    }
	for i := 0; i < 10; i++ {
        resultante[2*i]   = vetor1[i]  // Par
        resultante[2*i+1] = vetor2[i]  // Ímpar
    }
	for _, v := range resultante {
        fmt.Printf("%d ", v)
    }


	
}