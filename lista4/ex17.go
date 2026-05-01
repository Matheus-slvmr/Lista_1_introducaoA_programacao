package main
import "fmt"

func ehPrimo(n int) bool {
    if n <= 1 {
        return false
    }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}

func main() {
    var vetor [10]int
    
    fmt.Println("Digite 10 números:")
    for i := 0; i < 10; i++ {
        fmt.Printf("Número %d: ", i+1)
        fmt.Scan(&vetor[i])
    }
    
    fmt.Println("\nNúmeros primos:")
    for i := 0; i < 10; i++ {
        if ehPrimo(vetor[i]) {
            fmt.Printf("Número %d é primo na posição %d\n", vetor[i], i)
        }
    }
}