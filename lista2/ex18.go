package main
import (
	"fmt"
)
func main() {
	var preco_normal float64
	fmt.Print("Digite o preço normal do produto: ")
	fmt.Scan(&preco_normal)
	fmt.Print("Digite o tipo de produto (comum ou lançamento): ")
	var tipo string
	fmt.Scan(&tipo)
	var data int 
	fmt.Println("-----------datas--------")
	fmt.Println("segunda -feira (2)\n, terça-feira (3)\n, quarta-feira (4)\n, quinta-feira (5)\n , sexta-feira (6)\n, sadado (7) \n, domingo (1): ")
	fmt.Print("digite a data: ")
	fmt.Scan(&data)
	var preco_final float64
	if tipo == "comum" {
		switch data {
		case 1,4,6,7: // domingo, quarta-feira, sexta-feira e sábado preço normal
			preco_final = preco_normal
		case 2,3,5: // segunda-feira, terça-feira, quinta-feira 40%
			preco_final = float64(preco_normal) * 0.6 
		default:
			fmt.Println("Data inválida.")
			return
		}
	} else if tipo == "lançamento" {
		switch data {
		case 1,6,7,4: // domingo, sexta-feira, sábado e quarta-feira acrescimo de 15%
			preco_final = float64(preco_normal) * 1.15 
		case 2,3,5: // segunda-feira, terça-feira, quinta-feira desconto de 25%
			preco_final = float64(preco_normal) * 0.75 
		default:
			fmt.Println("Data inválida.")
			return
		}
	} else {
		fmt.Println("Tipo de produto inválido.")
		return
	}
	fmt.Printf("Preço final do produto: R$ %.2f\n", preco_final)
}