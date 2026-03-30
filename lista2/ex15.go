package main

import "fmt"

func main() {
	var dia, mes, ano int
	fmt.Scan(&dia, &mes, &ano)

	meses := []string{"", "janeiro", "fevereiro", "março", "abril", "maio", "junho", "julho", "agosto", "setembro", "outubro", "novembro", "dezembro"}

	if mes >= 1 && mes <= 12 {
		fmt.Printf("%d de %s de %d\n", dia, months[mes], ano)
	} else {
		fmt.Println("Mês inválido")
	}
}
