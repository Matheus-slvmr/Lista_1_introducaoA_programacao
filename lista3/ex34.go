package main

import f "fmt"

func main() {
	var n1, n2, mmc uint

	f.Println("Insira dois números :")
	f.Scan(&n1, &n2)
	if n1 <= 0 || n2 <= 0 {
		f.Println("Número inválido")
		return
	}
	if n2 > n1 {
		n1, n2 = n2, n1
	}
	for i := n2; i < n1*n1; i++ {
		if i%n2 == 0 && i%n1 == 0 {
			mmc = i
			break
		}
	}
	f.Printf("MMC(%d,%d)=%d", n1, n2, mmc)
}