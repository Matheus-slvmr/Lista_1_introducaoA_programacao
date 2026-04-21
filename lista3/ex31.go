package main

import f "fmt"

func main() {
	var total uint64 = 0
	var grao uint64 = 1

	for i := 1; i <= 64; i++ {
		total = total + grao
		grao = grao * 2
	}

	f.Println("O total de grãos é de: ", total)
}