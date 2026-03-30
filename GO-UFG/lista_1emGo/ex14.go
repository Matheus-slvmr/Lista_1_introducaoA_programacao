package main

import (
	"fmt"
	"math"
)

func main() {
	var altura float64
	var aresta float64
	fmt.Scan(&altura, &aresta)
	base := (3 * (aresta * aresta) * math.Sqrt(3)) / 2
	volumepiramide := (1.0 / 3.0) * base * altura
	fmt.Printf("O VOLUME DA PIRAMIDE E = %.2f\n METROS CUBICOS", volumepiramide)
}