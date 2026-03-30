package main
import (
	f "fmt"
)
func main (){
	num := 0
f.Print("digite um valor inteiro: ")
f.Scan(&num)
if num % 2 != 0 {
	f.Print("seu numero é impar")
}else{
	f.Print("seu numero é par")
}
}