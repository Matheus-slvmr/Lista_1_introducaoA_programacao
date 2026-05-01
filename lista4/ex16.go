package main
import "fmt"

func main() {
    var idades [50]int
    var frequencia [101]int // idades de 0 a 100
    var moda, maiorFreq int
    
    fmt.Println("Digite 50 idades:")
    for i := 0; i < 50; i++ {
        fmt.Printf("Idade %d: ", i+1)
        fmt.Scan(&idades[i])
        frequencia[idades[i]]++
    }
    
    maiorFreq = 0
    for i := 0; i <= 100; i++ {
        if frequencia[i] > maiorFreq {
            maiorFreq = frequencia[i]
            moda = i
        }
    }
    
    fmt.Printf("\nA moda é %d (apareceu %d vezes)\n", moda, maiorFreq)
}