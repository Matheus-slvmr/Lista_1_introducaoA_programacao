/* Escreva uma função em GO chamada MEDIA que retorne a média de 3 valores reais (X, Y e Z) passados como parâmetros. */
package main
import f "fmt"

func main() {
   x, y, z := 0.0, 0.0, 0.0
   f.Print("Digite o primeiro valor: ")
   f.Scan(&x)
   f.Print("Digite o segundo valor: ")
   f.Scan(&y)
   f.Print("Digite o terceiro valor: ")
   f.Scan(&z)
   m := media(x, y, z)
   f.Printf("A média entre %.2f, %.2f e %.2f é %.2f", x, y, z, m)
}

//  função media
func media(x, y, z float64) float64 {
   return (x + y + z) / 3
}