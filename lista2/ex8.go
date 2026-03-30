/*
Faça um programa que indique se um número inteiro informado pelo usuário está compreendido
entre 20 e 90 ou não. (20 e 90 não estão na faixa de valores).
*/
package main
import (
	f "fmt"
)
func main (){
	num1 := 0 
f.Print("digite um valores inteiros: ")
f.Scan(&num1)
if num1 > 20 && num1 < 90 {
	f.Printf("seu numero esta compreendido entre o 20 e o 90, seu numero é %v", num1)
}else {
	
	f.Printf("seu numero nao esta compreendido entre 20 e 90 , seu numero é %v", num1)
	 
	}
}