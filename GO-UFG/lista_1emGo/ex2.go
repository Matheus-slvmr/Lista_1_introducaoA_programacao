package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var resultados []float64
	for i := 0; i < n; i++ {
		var publico int
		var ppopular, pgeral, parquibanca, pcadeiras float64
		fmt.Scan(&publico, &ppopular, &pgeral, &parquibanca, &pcadeiras)
		vipopular := float64(publico) * (ppopular / 100)
		vigeral := float64(publico) * (pgeral / 100)
		viarquibanca := float64(publico) * (parquibanca / 100)
		vicadeiras := float64(publico) * (pcadeiras / 100)

		vf := (vipopular * 1) + (vigeral * 5) + (viarquibanca * 10) + (vicadeiras * 20)

		resultados = append(resultados, vf)
	}

	for i := 0; i < n; i++ {
		fmt.Printf("A RENDA DO JOGO N. %d E = %.2f\n", i+1, resultados[i])
	}
}