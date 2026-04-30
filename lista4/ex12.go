package main

import "fmt"

func main() {
    var vetorA [15]int  // Mudança: int ao invés de float64
    contagem := make(map[int]int)  // Mudança: map[int]int

    // 1. Leitura das 15 notas
    fmt.Println("Digite 15 notas (0 a 10):")
    for i := 0; i < 15; i++ {
        fmt.Scan(&vetorA[i])
        // Validação
        if vetorA[i] < 0 || vetorA[i] > 10 {
            vetorA[i] = 0
        }
        contagem[vetorA[i]]++
    }

    // 2. Imprime TABELA COMPLETA (todas as notas 0 a 10)
    fmt.Println("\n=== TABELA DE FREQUÊNCIAS ===")
    fmt.Printf("%-6s %-12s %-12s\n", "Nota", "Freq.Absoluta", "Freq.Relativa")
    fmt.Println("-------------------------------------")

    // Percorre TODAS as notas possíveis (0 a 10)
    for nota := 0; nota <= 10; nota++ {
        qtd := contagem[nota]  // Pega frequência absoluta
        relativa := float64(qtd) / 15 * 100  // Frequência relativa em %
        
        fmt.Printf("%-6d %-12d %.2f%%\n", nota, qtd, relativa)
    }
}