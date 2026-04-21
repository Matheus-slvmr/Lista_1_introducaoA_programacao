package main

import f "fmt"

func fatorial(n int) int {
	var P int = 1

	for i := 1; i <= n; i++ {
		P = P * i
	}
	return P
}
func main() {
	var x float64
	x = 100.0

	for i := 1; i <= 19; i++ {
		x = x + (100-float64(i))/float64(fatorial(i))
	}
	f.Println("O valor da soma dos 20 primeeiros termmos é: ", x)
}