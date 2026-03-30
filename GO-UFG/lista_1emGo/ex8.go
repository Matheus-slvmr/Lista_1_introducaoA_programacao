package main

import "fmt"

func main() {
	var raio float64
	var altura float64
	fmt.Scan(&raio, &altura)
	areacirculo := (3.14159 * (raio * raio))
	arealateral := (2 * 3.14159 * raio * altura)
	areatotal := arealateral + (2 * areacirculo)
	custo := areatotal * 100
	fmt.Printf("O VALOR DO CUSTO E = %.2f\n", custo)
}