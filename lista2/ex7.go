/*
Escreva um programa que leia três valores inteiros distintos (assuma que o usuário digitará valores
diferentes entre si) e os armazene nas variáveis A, B e C. Em seguida, descubra o menor valor e o
armazene em uma variável denominada MENOR; o maior valor, coloque-o na variável MAIOR e o
valor intermediário, na variável INTER. Imprima os valores em ordem crescente, ou seja, imprima
as variáveis MENOR, INTER e MAIOR, nessa ordem.
*/
package main
import (
	f "fmt"
)
func main (){
	a , b , c , menor , maior, inter := 0 , 0 , 0 , 0 , 0 , 0
f.Print("digite tres valores inteiros: ")
f.Scan(&a,&b,&c)

if a < b && a < c {
	menor = a
}else if (a < b && a > c) || (a < c && a > b){
	inter = a
}else {
	maior = a
}
if b < a && b < c {
	menor = b
}else if (b < a && b > c) || (b < c && b > a){
	inter = b
}else{
	maior = b
}
if c < a && c < b {
	menor = c
}else if (c < a && c > b ) || (c < b && c > a) {
	inter = c
}else{
	maior = c
}
f.Printf("em ordem crescente: %v %v %v", menor , inter , maior)



}


