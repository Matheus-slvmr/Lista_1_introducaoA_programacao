package main
import (
	f "fmt"
)
func main (){
	x := 0
	f.Print("digite um valor : ")
	f.Scan(&x)
	funcao := 8 / (2 - x)
	f.Printf("o resultado da f(x)=8 / (2-x)= %v", funcao)
}