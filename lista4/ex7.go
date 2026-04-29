package main

import "fmt"

func main() {
    var vetor [100]int

    // Leitura
    for i := 0; i < 200; i++ {
		if i%2 != 0 {
			vetor[i] = i 
		}else{
			continue
		}
		fmt.Print(vetor[i],"\n")
    }

	
}