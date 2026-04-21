package main

import f "fmt"

func main() {
	var n1, n2, produto float64

	f.Println("Insira dois números :")
	f.Scan(&n1, &n2)
	if n2 < 0 {
		n2, n1 = n1, n2
	}
	if n1 < 0 && n2 < 0 {
		n1 = (-1) * n1
		n2 = (-1) * n2
	}
	for i := 1.0; i <= n2; i++ {
		produto = produto + n1
	}
	f.Println("N1 x N2 = ", produto)
}