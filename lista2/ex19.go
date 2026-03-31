package main
import (
	"fmt"
	"math"
)
func main() {
	pi := math.Pi
	raio := 0.0
	altura := 0.0
	area := 0.0
	volume := 0.0
	tipo := ""
	fmt.Print("Digite o tipo de sólido (cilindro , cone ou esfera): ")
	fmt.Scan(&tipo)
	fmt.Print("Digite o raio da base: ")
	fmt.Scan(&raio)
	fmt.Print("Digite a altura: ")
	fmt.Scan(&altura)
	switch tipo {
	case "cilindro":
		area = 2 * pi * raio * altura
		volume = pi * (raio * raio) * altura
		fmt.Printf("a area do cilindro é: %v", area)
		fmt.Printf("o volume do cilindro é: %v", volume)
		case "cone":
		area = pi * raio * (math.Sqrt((raio*raio)+(altura*altura)))

		volume = (1.0 / 3.0) * pi * (raio * raio) * altura

		fmt.Printf("a area do cone é: %v", area)
		fmt.Printf("o volume do cone é: %v", volume)
		case "esfera":
		area = 4 * pi * (raio * raio)
		volume = (4.0 / 3.0) * pi * (raio * raio * raio)
		fmt.Printf("a area da esfera é: %v", area)
		fmt.Printf("o volume da esfera é: %v", volume)
	}
}

















}