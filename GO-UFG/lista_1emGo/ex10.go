package main

import "fmt"

func main() {
	var a float64
	var b float64
	var c float64
	var d float64
	fmt.Scan(&a, &b, &c, &d)
	detM := (a * d) - (b * c)
	fmt.Printf("O VALOR DO DETERMINANTE E = %.2f\n", detM)
}