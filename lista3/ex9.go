package main
import "fmt"
func main() {
	var nota1 , nota2 float64
	var alunos int
	countAprovados := 0
	countReprovados := 0
	couuntExame := 0
	somaClasse := 0.0
	fmt.Print("Digite a quantidade de alunos: ")
	fmt.Scan(&alunos)
	for i := 0; i < alunos; i++ {
		fmt.Printf("Digite a nota 1 do aluno %d: ", i+1)
		fmt.Scan(&nota1)
		fmt.Printf("Digite a nota 2 do aluno %d: ", i+1)
		fmt.Scan(&nota2)
		media := (nota1 + nota2) / 2
		somaClasse += media
		fmt.Printf("A média do aluno %d é: %.2f\n", i+1, media)
		if media >= 7 {
			fmt.Printf("O aluno %d foi aprovado.\n", i+1)
			countAprovados++
		} else if media >= 3 && media < 7 {
			fmt.Printf("O aluno %d está em exame.\n", i+1)
			couuntExame++
		} else {
			fmt.Printf("O aluno %d foi reprovado.\n", i+1)
			countReprovados++

		}
	}
	mediaClasse := somaClasse / float64(alunos)
	fmt.Printf("Quantidade de alunos aprovados: %d\n", countAprovados)
	fmt.Printf("Quantidade de alunos em exame: %d\n", couuntExame)
	fmt.Printf("Quantidade de alunos reprovados: %d\n", countReprovados)
	fmt.Printf("Média da classe: %.2f\n", mediaClasse)

}