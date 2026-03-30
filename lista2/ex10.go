package main
import (
	f "fmt"
)
func main (){
//constantes
	const regiaoNorteIda = 500
	const regiaoNorteIdaEVolta = 900
	const regiaoNordesteIda = 350
	const regiaoNordesteIdaEVolta = 650
	const regiaoCentroOesteIda = 350
	const regiaoCentroOesteIdaEVolta = 600
	const regiaoSulIda = 300
	const regiaoSulIdaEVolta = 550
//variavel
	num1 , num2 := 0 , 0

f.Print("digite 1 para regiao norte: ")
f.Print("digite 2 para regiao nordeste: ")
f.Print("digite 3 para regiao centro- oeste: ")
f.Print("digite 4 para regiao sul: ")

f.Scan(&num1)

switch num1 {
case 1: 
	f.Println("voce quer viagem de ida e volta?")
	f.Println("escreva 1 para sim, 2 para nao:")
	f.Scan(&num2) 
	if num2 == 2 {
	f.Println("o valor da sua viagem é R$:", regiaoNorteIda)
	}else if num2 == 1 {
	f.Println("o valor da sua viagem é R$:",regiaoNorteIdaEVolta)
	}
case 2:
	f.Println("voce quer viagem de ida e volta?")
	f.Println("escreva 1 para sim, 2 para nao:")
	f.Scan(&num2) 
	if num2 == 2 {
	f.Println("o valor da sua viagem é R$:", regiaoNordesteIda)
	}else if num2 == 1 {
	f.Println("o valor da sua viagem é R$:",regiaoNordesteIdaEVolta)
	}
case 3:
	f.Println("voce quer viagem de ida e volta?")
	f.Println("escreva 1 para sim, 2 para nao:")
	f.Scan(&num2) 
	if num2 == 2 {
	f.Println("o valor da sua viagem é R$:", regiaoCentroOesteIda)
	}else if num2 == 1 {
	f.Println("o valor da sua viagem é R$:",regiaoCentroOesteIdaEVolta)
	}
case 4:
	f.Println("voce quer viagem de ida e volta?")
	f.Println("escreva 1 para sim, 2 para nao:")
	f.Scan(&num2) 
	if num2 == 2 {
	f.Println("o valor da sua viagem é R$:", regiaoSulIda)
	}else if num2 == 1 {	f.Println("o valor da sua viagem é R$:",regiaoSulIdaEVolta)
	}
default:
	f.Print("Opção de região inválida")
}

}