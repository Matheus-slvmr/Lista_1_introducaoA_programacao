package main
import "fmt"

func main() {
	var base, expoente int
	fmt.Println("digite a base:")
	fmt.Scanln(&base)
	fmt.Print("digite o expoente: ")
	fmt.Scanln(&expoente)
	resultado := 1
	for i := 1; i <= expoente; i++ {
		resultado *= base
	}
	fmt.Println("resultado: ", resultado)
}