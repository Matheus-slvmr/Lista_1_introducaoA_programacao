package main

import "fmt"

func main() {
    var vetor [10]int
    var pares [10]int
    var impares [10]int
    var qtdPares, qtdImpares int
    var somaPares int = 0

    // 1. LEITURA dos 10 números
    fmt.Println("Digite 10 números inteiros:")
    for i := 0; i < 10; i++ {
        fmt.Printf("Número %d: ", i+1)
        fmt.Scan(&vetor[i])
    }

    // 2. PROCESSAMENTO: separa pares/ímpares e calcula soma
    for i := 0; i < 10; i++ {
        if vetor[i]%2 == 0 {        // É PAR?
            pares[qtdPares] = vetor[i]
            somaPares += vetor[i]    // Soma os pares
            qtdPares++
        } else {                     // É ÍMPAR
            impares[qtdImpares] = vetor[i]
            qtdImpares++
        }
    }

    // 3. SAÍDA formatada conforme exigido
    fmt.Println("\n=== RESULTADO ===")
    
    // a) NÚMEROS PARES
    fmt.Print("a) Números pares: [")
    for i := 0; i < qtdPares; i++ {
        fmt.Printf("%d", pares[i])
        if i < qtdPares-1 {
            fmt.Print(" ")
        }
    }
    fmt.Println("]")
    
    // b) SOMA DOS PARES
    fmt.Printf("b) Soma dos pares: %d\n", somaPares)
    
    // c) NÚMEROS ÍMPARES
    fmt.Print("c) Números ímpares: [")
    for i := 0; i < qtdImpares; i++ {
        fmt.Printf("%d", impares[i])
        if i < qtdImpares-1 {
            fmt.Print(" ")
        }
    }
    fmt.Println("]")
    
    // d) QUANTIDADE DE ÍMPARES
    fmt.Printf("d) Quantidade de ímpares: %d\n", qtdImpares)
}