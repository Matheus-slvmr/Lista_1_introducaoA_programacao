 package main
import (
	f "fmt"
)
func main (){
	num := 0
f.Print("digite um valor inteiro: ")
f.Scan(&num)
if num%5 != 0 {
	f.Print("seu numero nao é divisivel por 5")
}else{
	f.Print("seu numero é divisivel por 5")
}
}