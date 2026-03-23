package main

import "fmt"

func main()  {
	var lado_1 , lado_2 , lado_3 int
	
	fmt.Println("digite os lados do triangulo")

	fmt.Scan(&lado_1 , &lado_2 , &lado_3)

	if lado_1 == lado_2 && lado_1 == lado_3 {

		fmt.Println("seu triangulo é equilatero")
		
	} else if lado_1 == lado_2 || lado_1 == lado_3 || lado_2 == lado_3 || lado_3 == lado_2  {

fmt.Println("seu triangulo é isoceles")

	}else {

		fmt.Print("seu trinagulo é escaleno")
	}
}