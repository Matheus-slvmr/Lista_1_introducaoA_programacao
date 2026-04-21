package main

import f "fmt"

func main() {
	var n1, n2, quociente, resto, s int
	s = 0
	f.Println("Insira dois números :")
	f.Scan(&n1, &n2)
	if n1 < 0 || n2 < 0 {
		f.Println("Número inválido")
		return
	}
	for i := 1; i <= n1; i++ {
		s = s + n2
		if n1-s >= 0 {
			if n1-s < n2 {
				resto = n1 - s
				quociente = i
			}
		}
	}
	f.Printf("Quociente(%d,%d)=%d     Resto(%d,%d)=%d", n1, n2, quociente, n1, n2, resto)
}