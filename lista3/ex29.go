package main

import f "fmt"

func main() {
	var n int
	f.Println("Insira um número inteiro:")
	f.Scan(&n)

	if n < 0 {
		f.Println("Número inserido inválido")
		return
	}
	s := 0
	for i := 1; i <= n; i++ {
		s = s + i
	}

	f.Printf("A soma dos números inteiros de 1 até %d é %d", n, s)
}