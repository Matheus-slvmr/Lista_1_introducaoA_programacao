package main
import f "fmt"
func main (){

	inicial,num1,final := 0,0,0
	const ar, pintura, vidro, direcao = 1750, 800, 1200, 2000
	f.Scan(&inicial)
	f.Println("------------MENU------------")
	f.Println("1 = Ar condicionado")
	f.Println("2 = Pintura Metalica")
	f.Println("3 = Vidro eletrico")
	f.Println("4 = Direçao hidraulica")
	f.Scan(&num1)

	switch num1 {

case 1: 
    final = inicial + ar
	f.Println("Voce escolheu ar condicionado")
	f.Println("o preço final do seu carro será: %v",final)
case 2:
	final = inicial + pintura
	f.Println("Voce Pintura metalica")
	f.Println("o preço final do seu carro será: %v",final)
case 3:
	final = inicial + vidro
	f.Println("Voce escolheu Vidro eletrico")
	f.Println("o preço final do seu carro será: %v",final)
case 4:
	final = inicial + direcao
	f.Println("Voce escolheu Direçao hidráulica")
	f.Printf("o preço final do seu carro será: %v",final)
default:
	f.Print("Opção inválida")
}


}