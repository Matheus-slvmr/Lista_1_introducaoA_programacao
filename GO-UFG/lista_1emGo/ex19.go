package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	resultado := 0.0
	if n > 1 {
		for i := 1; i <= n; i++ {
			resultado += 1 / float64(i)
		}
		fmt.Printf("%.6f\n", resultado)
	} else {
		fmt.Println("Numero invalido!")
	}

}