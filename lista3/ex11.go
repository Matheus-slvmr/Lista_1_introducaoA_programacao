package main

import f "fmt"

func fatoriall(num int) int {
	var resultado int = 1

	for i := 1; i <= num; i++ {
		resultado *= i
	}
	
	return resultado 
}

func main() {
	var n int
	f.Print("Digite um número: ")
	f.Scan(&n)

	f.Print("O fatorial é: ", fatoriall(n))
}