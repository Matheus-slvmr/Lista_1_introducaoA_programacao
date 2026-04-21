package main

import f "fmt"

func main() {
	var lucro, p float64
	var qtd int

	p = 6.00
	qtd = 130
	maximo := lucro
	for p > 1 {

		lucro = float64(qtd)*p - 300
		if lucro > maximo {
			maximo = lucro
		}
		f.Printf("PREÇO:R$ %.2f    INGRESSOS:%.0d     LUCRO:%.2f", p, qtd, lucro)
		f.Println()
		p = p - 0.60
		qtd = qtd + 30

	}
	f.Printf("LUCRO MÁXIMO: R$%.2f", maximo)

}