package main
import "fmt"

func main() {
    var vetor [10]int
    
    fmt.Println("Digite 10 números em ordem crescente:")
    for i := 0; i < 10; i++ {
        fmt.Printf("Número %d: ", i+1)
        fmt.Scan(&vetor[i])
    }
    
    fmt.Println("\nVetor ordenado:")
    for i := 0; i < 10; i++ {
        fmt.Printf("%d ", vetor[i])
    }
    fmt.Println()
}