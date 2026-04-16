/* Escreva um programa que receba vários números inteiros e verifique se eles são ou não quadrados perfeitos. O
programa deve terminar quando o usuário informar um número menor ou igual a zero. Obs.: Um número é
quadrado perfeito quando tem um número inteiro como raiz quadrada. Não é permitido usar o comando sqrt
em sua solução. */
package main
import f "fmt"
func main(){
var numero int
for {
	f.Print("Digite um número inteiro (ou um número menor ou igual a zero para sair): ")
	f.Scan(&numero)	
	if numero <= 0 {
		break
	}
	// Verifica se o número é um quadrado perfeito
	quadradoperfeito := false
	for i := 1; i*i <= numero; i++ {
		if i*i == numero {
			quadradoperfeito = true
			break
		}
	}
	if quadradoperfeito {
		f.Printf("%d é um quadrado perfeito.\n", numero)
	} else {
		f.Printf("%d não é um quadrado perfeito.\n", numero)
	}
}
}