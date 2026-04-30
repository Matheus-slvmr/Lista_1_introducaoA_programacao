package main

import "fmt"

func main() {
    var vetor [50]int

a := 0
b := 1

vetor[0] = 0
vetor[1] = 1

    for i := 2; i < 50; i++ {
        proximo := a + b 
        a = b 
        b = proximo 
        vetor[i] = proximo
        
	}
    for i := 0; i < 50; i++ {
        fmt.Print(vetor[i],"\n")
    }
}
