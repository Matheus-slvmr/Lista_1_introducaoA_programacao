package main

import "fmt"

func main() {
	var nota float64
	fmt.Scan(&nota)
	fmt.Printf("NOTA = %.2f\n", nota)

	switch {
	case nota >= 9.0 && nota <= 10.0:
		fmt.Println("CONCEITO = A")
	case nota >= 7.5 && nota < 9.0:
		fmt.Println("CONCEITO = B")
	case nota >= 6.0 && nota < 7.5:
		fmt.Println("CONCEITO = C")
	case nota >= 0.0 && nota < 6.0:
		fmt.Println("CONCEITO = D")
	default:
		fmt.Println("Nota invalida")
	}
}