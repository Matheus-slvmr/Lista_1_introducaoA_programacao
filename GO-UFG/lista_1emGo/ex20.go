package main

import "fmt"

func main() {
	var horas int
	var minutos int
	var segundos int
	fmt.Scan(&horas, &minutos, &segundos)
	totalSegundos := (horas * 3600) + (minutos * 60) + segundos
	fmt.Printf("O TEMPO EM SEGUNDOS E = %d\n", totalSegundos)
}