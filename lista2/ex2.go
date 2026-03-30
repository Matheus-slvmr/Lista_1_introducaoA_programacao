package main
import (
	f "fmt"
)
func main (){
	num := 0
f.Print("digite um valor inteiro: ")
f.Scan(&num)
if num > 0 {
	f.Print("seu numero é positivo")
}else if num < 0 {
	f.Print("seu numero é negativo")
}else{
	f.Print("seu numero é nulo")
}
}