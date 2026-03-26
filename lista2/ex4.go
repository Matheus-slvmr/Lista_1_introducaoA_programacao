package main
import (
	f "fmt"
	m "math"
)
func main (){
	num := 0.0
f.Print("digite um valor : ")
f.Scan(&num)
if num >= 0 {
	num = m.Sqrt(num) 
	f.Print("sua raiz é",num)
}else {
	num = m.Pow(num , 2)
	f.Print("seu quadrado é ", num)
}
}