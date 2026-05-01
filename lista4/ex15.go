package main
import "fmt"

func main() {
    var vetor [30]int
    var novoVetor [30]int
    
    fmt.Println("Digite 30 números:")
    for i := 0; i < 30; i++ {
        fmt.Printf("Número %d: ", i+1)
        fmt.Scan(&vetor[i])
    }
    
    for i := 0; i < 30; i++ {
        if i%2 == 0 {
            novoVetor[i] = vetor[i] * 2 // posições pares = dobro
        } else {
            novoVetor[i] = vetor[i] * 3 // posições ímpares = triplo
        }
    }
    
    fmt.Println("\nNovo vetor:")
    for i := 0; i < 30; i++ {
        fmt.Printf("Posição %d: %d\n", i, novoVetor[i])
    }
}