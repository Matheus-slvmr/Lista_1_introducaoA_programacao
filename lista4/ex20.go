package main
import "fmt"

func main() {
    var jogadas [20]int
    var freq [7]int // 1 a 6
    
    fmt.Println("Digite 20 números sorteados (1 a 6):")
    for i := 0; i < 20; i++ {
        fmt.Printf("Jogada %d: ", i+1)
        fmt.Scan(&jogadas[i])
        freq[jogadas[i]]++
    }
    
    fmt.Println("\nNúmeros sorteados:", jogadas)
    fmt.Println("Frequência:")
    for i := 1; i <= 6; i++ {
        fmt.Printf("%d: %d vezes\n", i, freq[i])
    }
}