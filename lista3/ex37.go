/*
 37. Faça um programa que leia um número inteiro positivo na base 8, calcule e imprima o seu equivalente na base

10.
*/
package main

import (
	f "fmt"
	m "math"
	"strconv"
)

func main() {
	var n string
	var base int

	base = 0

	f.Println("Insira um número na base 8 :")
	f.Scan(&n)

	for i := 0; i < len(n); i++ {
		digito, err := strconv.Atoi(string(n[i]))
		if err != nil || digito < 0 || digito > 7 {
			f.Println("Digito inserido inválido")
			return
		}

		base = base + digito*int(m.Pow(8, float64(len(n)-i-1)))
	}

	f.Printf("O número %s na base 8 é igual a %d na base 10", n, base)
}