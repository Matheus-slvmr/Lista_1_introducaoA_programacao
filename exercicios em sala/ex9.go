package main

import "fmt"

type Criatura struct {
	Nome string
	Idade int
	Peso float64
	Altura float64
}

func main() {
	c := Criatura{
		Nome: "Jac and friends",
		Idade: 10,
		Peso: 50.5,
		Altura: 1.5,
	}
 	
	fmt.Println(c.Nome)
	fmt.Println(c.Idade)
	fmt.Println(c.Peso)
	fmt.Println(c.Altura)
}
