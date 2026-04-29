package main

import "fmt"

func main() {
    var vetor1 [10]int      // Primeiro vetor (10 elementos)
    var vetor2 [5]int       // Segundo vetor (5 elementos)
    
    // Leitura do primeiro vetor (10 elementos)
    fmt.Println("Digite 10 números para o primeiro vetor:")
    for i := 0; i < 10; i++ {
        fmt.Printf("Índice %d: ", i)
        fmt.Scan(&vetor1[i])
    }
    
    // Leitura do segundo vetor (5 elementos)
    fmt.Println("\nDigite 5 números para o segundo vetor:")
    for i := 0; i < 5; i++ {
        fmt.Printf("Índice %d: ", i)
        fmt.Scan(&vetor2[i])
    }
    
    // Calcula a soma de todos os elementos do segundo vetor
    somaVetor2 := 0
    for i := 0; i < 5; i++ {
        somaVetor2 += vetor2[i]
    }
    
    // Vetores resultantes
    var pares [10]int
    var impares [10]int
    var qtdPares, qtdImpares int
    
    // Separa pares e ímpares do primeiro vetor
    for i := 0; i < 10; i++ {
        if vetor1[i]%2 == 0 {
            pares[qtdPares] = vetor1[i] + somaVetor2
            qtdPares++
        } else {
            impares[qtdImpares] = vetor1[i] + somaVetor2
            qtdImpares++
        }
    }
    
    // Exibe os vetores
    fmt.Println("\n=== RESULTADO ===")
    fmt.Print("Vetor 1: [")
    for i := 0; i < 10; i++ {
        fmt.Printf("%d", vetor1[i])
        if i < 9 {
            fmt.Print(" ")
        }
    }
    fmt.Println("]")
    
    fmt.Print("Vetor 2: [")
    for i := 0; i < 5; i++ {
        fmt.Printf("%d", vetor2[i])
        if i < 4 {
            fmt.Print(" ")
        }
    }
    fmt.Println("]")
    
    fmt.Print("Vetor pares: [")
    for i := 0; i < qtdPares; i++ {
        fmt.Printf("%d", pares[i])
        if i < qtdPares-1 {
            fmt.Print(" ")
        }
    }
    fmt.Println("]")
    
    fmt.Print("Vetor ímpares: [")
    for i := 0; i < qtdImpares; i++ {
        fmt.Printf("%d", impares[i])
        if i < qtdImpares-1 {
            fmt.Print(" ")
        }
    }
    fmt.Println("]")
}