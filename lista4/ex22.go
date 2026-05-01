package main
import "fmt"

func main() {
    var codigos [10]int
    var saldos [10]float64
    var nContas int
    var opcao, codigo, pos int
    var valor float64
    
    // Cadastro inicial
    fmt.Print("Quantas contas cadastrar (max 10)? ")
    fmt.Scan(&nContas)
    
    for i := 0; i < nContas; i++ {
        fmt.Printf("Código da conta %d: ", i+1)
        fmt.Scan(&codigos[i])
        fmt.Printf("Saldo da conta %d: ", codigos[i])
        fmt.Scan(&saldos[i])
    }
    
    for {
        fmt.Println("\n=== MENU BANCO ===")
        fmt.Println("1 - Depósito")
        fmt.Println("2 - Saque")
        fmt.Println("3 - Ativo bancário")
        fmt.Println("4 - Finalizar")
        fmt.Print("Opção: ")
        fmt.Scan(&opcao)
        
        if opcao == 4 {
            fmt.Println("Programa finalizado!")
            break
        } else if opcao == 1 {
            fmt.Print("Código da conta: ")
            fmt.Scan(&codigo)
            pos = -1
            for i := 0; i < nContas; i++ {
                if codigos[i] == codigo {
                    pos = i
                    break
                }
            }
            if pos == -1 {
                fmt.Println("Conta não encontrada!")
                continue
            }
            fmt.Print("Valor do depósito: ")
            fmt.Scan(&valor)
            saldos[pos] += valor
            fmt.Printf("Novo saldo: R$ %.2f\n", saldos[pos])
            
        } else if opcao == 2 {
            fmt.Print("Código da conta: ")
            fmt.Scan(&codigo)
            pos = -1
            for i := 0; i < nContas; i++ {
                if codigos[i] == codigo {
                    pos = i
                    break
                }
            }
            if pos == -1 {
                fmt.Println("Conta não encontrada!")
                continue
            }
            fmt.Print("Valor do saque: ")
            fmt.Scan(&valor)
            if saldos[pos] >= valor {
                saldos[pos] -= valor
                fmt.Printf("Novo saldo: R$ %.2f\n", saldos[pos])
            } else {
                fmt.Println("Saldo insuficiente!")
            }
            
        } else if opcao == 3 {
            var total float64
            for i := 0; i < nContas; i++ {
                total += saldos[i]
            }
            fmt.Printf("Ativo bancário: R$ %.2f\n", total)
        }
    }
}