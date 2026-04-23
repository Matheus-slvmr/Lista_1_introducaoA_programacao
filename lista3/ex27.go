package main

import (
	f "fmt"
	m "math"
)

func fatorial(n int) int {
	var P int = 1

	for i := 1; i <= n; i++ {
		P = P * i
	}
	return P
}

func main() {
	var x float64
	f.Println("Insira o valor de x: ")
	f.Scan(&x)

	y := 1.0

	for i := 1; i <= 19; i++ {
		y = y + float64(m.Pow(x, float64(i+1)))*float64(m.Pow(-1, float64(i))/float64(fatorial(2*i)))
	}
	f.Println("Cosseno(x)=", y)
}