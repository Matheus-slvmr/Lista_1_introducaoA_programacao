package main

import "fmt"

func main() {
	var cc int
	var mc float64
	var RCI string
	fmt.Scan(&cc, &mc, &RCI)

	var vfc float64
	switch RCI {
	case "R":
		vfc = 5 + (mc * 0.05)
	case "C":
		vfc = 500 + ((mc - 80) * 0.25)
	default:
		vfc = 800 + ((mc - 100) * 0.04)
	}
	fmt.Printf("CONTA = %.d\n", cc)
	fmt.Printf("VALOR DA CONTA = R$%.2f\n", vfc)
}