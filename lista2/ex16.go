/*
Desenvolver um programa que leia os coeficientes (A, B e C) de uma equação do segundo grau ( Ax2 + Bx + C
=0) e que calcule suas raízes. O programa deve mostrar, quando possível, o valor das raízes calculadas e a
classificação das mesmas: “RAÍZES IMAGINÁRIAS”, “RAIZ ÚNICA” ou “RAÍZES DISTINTAS”.
*/
package main
import (
	"fmt"
	"math"
)
func main() {
	var a, b, c float64
	fmt.Print("Digite o valor de A: ")
	fmt.Scan(&a)
	fmt.Print("Digite o valor de B: ")
	fmt.Scan(&b)
	fmt.Print("Digite o valor de C: ")
	fmt.Scan(&c)
	delta := b*b - 4*a*c
	if delta < 0 {
		fmt.Println("RAÍZES IMAGINÁRIAS")
	} else if delta == 0 {
		raiz := -b / (2 * a)
		fmt.Printf("RAIZ ÚNICA: %.2f\n", raiz)
	} else {
		raiz1 := (-b + math.Sqrt(delta)) / (2 * a)
		raiz2 := (-b - math.Sqrt(delta)) / (2 * a)
		fmt.Printf("RAÍZES DISTINTAS: %.2f e %.2f\n", raiz1, raiz2)
	}
}