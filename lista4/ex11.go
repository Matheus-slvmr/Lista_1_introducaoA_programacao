
 package main
 import "fmt"
 func main (){
	var vetor[100] int
	var sum int
	for i := 0; i < 100; i++ {
		fmt.Scan(&vetor[i])
	}
	for i := 0; i <= 49; i++ {
        diferenca := vetor[i] - vetor[99-i]
        sum += diferenca * diferenca * diferenca  
    }
	fmt.Print(sum)
 }