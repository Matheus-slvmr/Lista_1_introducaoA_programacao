package main

import f "fmt"

func main() {
	var n, n1 int
	var binario string = ""

	f.Println("Insira um número na base decimal :")
	f.Scan(&n)
	n1 = n
	if n < 0 {
		f.Println("Número inválido")
		return
	}
	if n == 0 {
		binario = "0"
	}
	for n > 0 {
		resto := n % 2
		if resto == 0 {
			binario = "0" + binario
		} else {
			binario = "1" + binario
		}
		n = n / 2
	}
	f.Printf(" %d na base 2 é: %s", n1, binario)
}