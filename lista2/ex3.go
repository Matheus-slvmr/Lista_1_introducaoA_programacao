package main
import (
	f "fmt"
)
func main (){
	num1 , num2, soma := 0 , 0 , 0
f.Print("digite um valor inteiro: ")
f.Scan(&num1 , &num2)
soma = (num1 + num2)
if soma > 20 {
soma = soma + 8
f.Println("seu valor é ", soma)
}else{
	soma = soma - 5
	f.Println("seu valor é ", soma)
}
}