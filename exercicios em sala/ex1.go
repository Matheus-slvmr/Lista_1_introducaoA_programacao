//30/03
package main
import (
	"fmt"
)
func main() {
	const numNotas int = 3 
	var(
		nota[numNotas]float64
	
	)
	for i := 0; i < numNotas; i++ {
		fmt.Printf("Digite a nota %d: ", i)
		fmt.Scan(&nota[i])
	    
	
	if nota[i] >= 6 {
		fmt.Printf("Aluno aprovado com nota %.2f\n", nota[i])
	} else {
		fmt.Printf("Aluno reprovado com nota %.2f\n", nota[i])
	}
}
}