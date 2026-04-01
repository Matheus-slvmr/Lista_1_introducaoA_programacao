package main
import f "fmt"
func main() {
   // Declaração e leitura dos números
   n:= 0
   f.Print("Digite um número inteiro: ")
   f.Scan(&n)
   // Chamada à função fatorial
   fatorial := fatorial(n)
   f.Printf("O fatorial de %d é %d", n, fatorial)
}
func fatorial(numero int) int {
   if numero == 0 || numero == 1 {
	  return 1
   }
   return numero * fatorial (numero - 1)
}
