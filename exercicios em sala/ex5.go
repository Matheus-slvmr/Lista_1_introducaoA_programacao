/* Escreva uma função em GO que leia 3 números inteiros e retorne o maior deles. */
package main
import f "fmt"
func main() {
   // Declaração e leitura dos números
   n1, n2, n3 := 0, 0, 0
   f.Print("Digite o primeiro número: ")
   f.Scan(&n1)
   f.Print("Digite o segundo número: ")
   f.Scan(&n2)
   f.Print("Digite o terceiro número: ")
   f.Scan(&n3)
   // Chamada à função
   m := maior(n1, n2, n3)
   f.Printf("O maior número entre %d, %d e %d é %d", n1, n2, n3, m)
}
// Definição da função maior
func maior(numero1, numero2, numero3 int) int {
   if numero1 >= numero2 && numero1 >= numero3 {
      return numero1
   } else if numero2 >= numero1 && numero2 >= numero3 {
      return numero2
   } else {
      return numero3
   }
}