package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	if x%2 == 0 {
		for i := 0; i < y; i++ {
			fmt.Printf("%d\n", i+x+i)
		}
	} else {
		fmt.Println("O PRIMEIRO NUMERO NAO E PAR")
	}
}