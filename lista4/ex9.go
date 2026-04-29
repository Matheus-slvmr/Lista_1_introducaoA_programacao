package main
import "fmt"
func main (){
	var altura[10]float64
	var media float64
	
	for i := 0; i < 10; i++ {
		fmt.Scan(&altura[i])
	}
	soma := 0.0
	for i := 0; i < 10; i++ {
		soma += altura[i]
	}
	media = soma / 10.0
	fmt.Println("media",media)
	for i := 0; i < 10; i++ {
        if altura[i] > media {
            fmt.Printf("   Atleta %d: %.2f m\n", i, altura[i])
        }
    }
}