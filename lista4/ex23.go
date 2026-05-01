package main
import "fmt"

func main() {
    var janela [24]int
    var corredor [24]int
    var tipo int // 1=janela, 2=corredor
    var poltrona, livresJanela, livresCorredor int
    
    // Inicializa poltronas vazias
    for i := 0; i < 24; i++ {
        janela[i] = 0
        corredor[i] = 0
    }
    
    for {
        fmt.Println("\n=== VENDA DE PASSAGENS ===")
        fmt.Println("1 - Janela")
        fmt.Println("2 - Corredor")
        fmt.Println("0 - Sair")
        fmt.Print("Escolha: ")
        fmt.Scan(&tipo)
        
        if tipo == 0 {
            break
        }
        
        // Conta poltronas livres
        livresJanela = 0
        livresCorredor = 0
        for i := 0; i < 24; i++ {
            if janela[i] == 0 {
                livresJanela++
            }
            if corredor[i] == 0 {
                livresCorredor++
            }
        }
        
        if tipo == 1 {
            if livresJanela == 0 {
                fmt.Println("Não há poltronas livres na janela!")
                continue
            }
            fmt.Printf("Poltronas livres na janela: %d\n", livresJanela)
            fmt.Print("Número da poltrona (0-23): ")
            fmt.Scan(&poltrona)
            if poltrona >= 0 && poltrona < 24 && janela[poltrona] == 0 {
                janela[poltrona] = 1
                fmt.Println("Poltrona reservada!")
            } else {
                fmt.Println("Poltrona inválida ou ocupada!")
            }
        } else if tipo == 2 {
            if livresCorredor == 0 {
                fmt.Println("Não há poltronas livres no corredor!")
                continue
            }
            fmt.Printf("Poltronas livres no corredor: %d\n", livresCorredor)
            fmt.Print("Número da poltrona (0-23): ")
            fmt.Scan(&poltrona)
            if poltrona >= 0 && poltrona < 24 && corredor[poltrona] == 0 {
                corredor[poltrona] = 1
                fmt.Println("Poltrona reservada!")
            } else {
                fmt.Println("Poltrona inválida ou ocupada!")
            }
        }
        
        // Verifica se ônibus lotado
        if livresJanela == 0 && livresCorredor == 0 {
            fmt.Println("ÔNIBUS LOTADO!")
            break
        }
    }
    
    fmt.Println("\nEstado final das poltronas:")
    fmt.Print("Janela:  ")
    for i := 0; i < 24; i++ {
        fmt.Printf("%d ", janela[i])
    }
    fmt.Println()
    fmt.Print("Corredor:")
    for i := 0; i < 24; i++ {
        fmt.Printf("%d ", corredor[i])
    }
    fmt.Println()
}