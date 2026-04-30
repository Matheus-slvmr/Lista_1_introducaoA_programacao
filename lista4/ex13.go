package main

import "fmt"

func main() {
    var numEmpregado, mesesTrabalho int
    var dados [100][2]int  // [numEmpregado][mesesTrabalho]
    var n int = 0  // Contador de empregados lidos
    
    fmt.Println("Digite os dados dos empregados (0 0 para terminar):")
    
    // 1. Leitura dos dados até encontrar 0 0
    for {
        fmt.Scan(&numEmpregado, &mesesTrabalho)
        
        // Condição de parada
        if numEmpregado == 0 && mesesTrabalho == 0 {
            break
        }
        
        // Armazena os dados
        dados[n][0] = numEmpregado
        dados[n][1] = mesesTrabalho
        n++
    }
    
    // 2. Encontra os 3 com MENOS meses (mais recentes)
    var maisRecente1, maisRecente2, maisRecente3 [2]int
    maisRecente1[1] = 1000  // Inicializa com valor alto
    maisRecente2[1] = 1000
    maisRecente3[1] = 1000
    
    for i := 0; i < n; i++ {
        meses := dados[i][1]
        
        if meses < maisRecente1[1] {
            // Desloca posições
            maisRecente3 = maisRecente2
            maisRecente2 = maisRecente1
            maisRecente1[0] = dados[i][0]
            maisRecente1[1] = meses
        } else if meses < maisRecente2[1] {
            maisRecente3 = maisRecente2
            maisRecente2[0] = dados[i][0]
            maisRecente2[1] = meses
        } else if meses < maisRecente3[1] {
            maisRecente3[0] = dados[i][0]
            maisRecente3[1] = meses
        }
    }
    
    // 3. Imprime os 3 mais recentes
    fmt.Println("\n=== TRÊS EMPREGADOS MAIS RECENTES ===")
    fmt.Printf("1º: Empregado %d (%d meses)\n", 
               maisRecente1[0], maisRecente1[1])
    fmt.Printf("2º: Empregado %d (%d meses)\n", 
               maisRecente2[0], maisRecente2[1])
    fmt.Printf("3º: Empregado %d (%d meses)\n", 
               maisRecente3[0], maisRecente3[1])
}