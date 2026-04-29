package main

import "fmt"
import "math"

func main() {
    var vetor [15]float64
	var numeros int

    // Leitura
    for i := 0; i < 15; i++ {
		fmt.Scan(&numeros)
		if numeros < 0 {
			vetor[i] = -1
		}else{
			vetor[i] = math.Sqrt(float64(numeros))
		}
    }
	for i := 0; i < 15; i++ {
        fmt.Printf("vetor[%d] = %.2f\n", i, vetor[i])
    }

	
}
