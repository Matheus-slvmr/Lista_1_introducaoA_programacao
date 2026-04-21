package main

import f "fmt"

func main() {
	var peso, maiorpeso, menorpeso float64
	var numero, qtd, idmaior, idmenor int
	f.Println("Quantidade de bois:")
	f.Scan(&qtd)
	maiorpeso = 0

	for i := 1; i <= qtd; i++ {
		f.Println("Número identificador:")
		f.Scan(&numero)
		f.Println("Peso do boi")
		f.Scan(&peso)
		if peso > maiorpeso {
			maiorpeso = peso
			idmaior = numero
		}
		if i == 1 {
			menorpeso = peso
			idmenor = numero
		}
	}
	f.Printf("Boi mais gordo:   PESO = %f    ID=%d", maiorpeso, idmaior)
	f.Println()
	f.Printf("Boi mais gordo:   PESO = %f    ID=%d", menorpeso, idmenor)
}