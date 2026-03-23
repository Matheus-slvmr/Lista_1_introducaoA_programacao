package main

import "fmt"

func main()  {
var(
	nota_portugues float32
	nota_matematica float32
	peso_portugues float32 = 2
	peso_matematica float32 = 4
	media float32
)
fmt.Println("escreva sua nota de portugues:")
fmt.Scan(&nota_portugues)
fmt.Println("escreva sua nota de matematica:")
fmt.Scan(&nota_matematica)
media = ((nota_portugues * peso_portugues) + (nota_matematica * peso_matematica)) / (peso_matematica + peso_portugues)
fmt.Println("sua media ponderada é :", media)

}
