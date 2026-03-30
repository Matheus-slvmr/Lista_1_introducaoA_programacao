package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	taxaa := 10 * (a / 3)
	taxab := 5 * (a % 3)
	taxaf := float64(taxaa) + float64(taxab)
	fmt.Printf("O VALOR A PAGAR E = %.2f", taxaf)
}