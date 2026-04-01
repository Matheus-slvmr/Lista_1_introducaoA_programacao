package main
import f "fmt"
func main() {
   // Declaração e leitura dos números
   n:= 0
   f.Print("Digite um número inteiro: ")
   f.Scan(&n)
   // Chamada à função fatorial
   fatorial := fatorialIterativo(n)
   f.Printf("O fatorial de %d é %d", n, fatorial)
}
func fatorialIterativo(numero int) int {
    resultado := 1
    for i := 2; i <= numero; i++ {
        resultado *= i 
    }
    return resultado
}
