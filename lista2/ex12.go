package main
import (
	f "fmt"
)
func main (){
	idade := 0 
f.Print("digite sua idade: ")
f.Scan(&idade)
if idade <= 2 && idade >= 0 {
	f.Print("Recem-nascido")
}else if idade <= 11 && idade >= 3{
	f.Print("Criança") 
}else if idade >= 12 && idade <= 19{
		f.Print("Adolescente")
}else if idade <= 55 && idade >= 20{
		f.Print("Adulto")
}else if idade > 55{
		f.Print("idoso")
}
}