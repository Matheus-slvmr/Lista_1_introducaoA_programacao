package main

import "fmt"

func main() {
	var fahrenheit float64
	var chuva float64
	fmt.Scan(&fahrenheit, &chuva)

	celcius := (fahrenheit - 32) * 5 / 9
	chuvaMM := chuva * 25.4

	fmt.Printf("O VALOR EM CELSIUS = %.2f\n", celcius)
	fmt.Printf("A QUANTIDADE DE CHUVA E = %.2f\n", chuvaMM)
}