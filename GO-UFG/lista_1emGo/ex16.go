package main

import "fmt"

func main() {
	var salario float64
	fmt.Scan(&salario)
	switch {
	case salario <= 300.00:
		fmt.Printf("SALARIO COM REAJUSTE = %.2f\n", salario*1.5)

	default:
		fmt.Printf("SALARIO COM REAJUSTE = %.2f\n", salario*1.3)
	}
}