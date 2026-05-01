package main
import "fmt"

func main() {
    var vetor [10]float64
    var codigo int
    
    fmt.Println("Digite 10 números reais:")
    for i := 0; i < 10; i++ {
        fmt.Printf("Vetor[%d]: ", i)
        fmt.Scan(&vetor[i])
    }
    
    fmt.Print("Digite o código (0=sair, 1= direto, 2=inverso): ")
    fmt.Scan(&codigo)
    
    if codigo == 0 {
        fmt.Println("Programa finalizado")
    } else if codigo == 1 {
        fmt.Println("Ordem direta:")
        for i := 0; i < 10; i++ {
            fmt.Printf("%.2f ", vetor[i])
        }
        fmt.Println()
    } else if codigo == 2 {
        fmt.Println("Ordem inversa:")
        for i := 9; i >= 0; i-- {
            fmt.Printf("%.2f ", vetor[i])
        }
        fmt.Println()
    }
}