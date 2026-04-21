package main

import (
	"fmt"
)
func main (){
var h float64
 for i := 1.0; i <= 50.0; i++ {
		
		// Calcula a parte de cima
		cima := (i * 2.0) - 1.0 
		
		// A parte de baixo é i
		divisor := i            

		termoAtual := cima / divisor
		h += termoAtual
		fmt.Printf("Termo %.0f: %.0f / %.0f = %.4f\n", i, cima, divisor, termoAtual)
 }
	

	fmt.Printf("\nO valor final calculado para H é: %.10f\n", h)
}