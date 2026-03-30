package main
import (
	f "fmt"
)
func main (){
	num1 , num2 := 0 , 0
f.Print("digite dois valores inteiros: ")
f.Scan(&num1, &num2)
divisao := num1 / num2
if num1 % num2 != 0 {
	f.Print("seu numero nao é divisivel")
}else {
	
	f.Printf("seu numero %v é divisivel por %v que é %v", num1 , num2,divisao)
	 
}
}