package main
import (
	"fmt"
)
func main() {
	const numNotas int = 3 
	var(
		nota[numNotas]float64
		soma float64
	
	)
	for i := 0; i < numNotas; i++ {
		fmt.Printf("Digite a nota %d: ", i)
		fmt.Scan(&nota[i])
		soma += nota[i]
	}
	for i,v:= range nota {
		fmt.Printf("Nota %d: %.2f\n", i, v)
	}
	fmt.Printf("Soma das notas: %.2f\n", soma)
}