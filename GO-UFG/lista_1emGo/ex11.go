package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a%5 == 0 && a%3 == 0 {
		fmt.Println("O NUMERO E DIVISIVEL")
	} else {
		fmt.Println("O NUMERO NAO E DIVISIVEL")
	}
}