package main
import "fmt"

func main() {
    var num [10]int
    var divis [5]int
    
    fmt.Println("Digite 10 números para o vetor principal:")
    for i := 0; i < 10; i++ {
        fmt.Printf("Num[%d]: ", i)
        fmt.Scan(&num[i])
    }
    
    fmt.Println("Digite 5 números para os divisores:")
    for i := 0; i < 5; i++ {
        fmt.Printf("Divis[%d]: ", i)
        fmt.Scan(&divis[i])
    }
    
    for i := 0; i < 10; i++ {
        fmt.Printf("\nNúmero %d:\n", num[i])
        for j := 0; j < 5; j++ {
            if num[i]%divis[j] == 0 {
                fmt.Printf("Divisível por %d na posição %d\n", divis[j], j)
            }
        }
    }
}