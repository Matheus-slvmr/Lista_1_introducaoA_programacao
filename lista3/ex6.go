/* Gabriel inventou um código para representar números naturais, usando uma sequência de zeros e uns. Funciona assim, o número natural é representado pela quantidade de vezes que o padrão "100" aparece na sequência.

Por exemplo, na sequência 11101001010011110, o padrão aparece duas vezes e na sequência 11101010111110111010101 ele não aparece nenhuma vez. Você deve ajudar Gabriel implementar um programa que, dada a sequência de zeros e uns, calcule quantas vezes o padrão "100" aparece nela.

Entrada
A primeira linha da entrada contém um inteiro N, o tamanho da sequência. A segunda linha contém a sequência de N zeros e uns, separados por espaço em branco.

Saída
Seu programa deve imprimir um inteiro, quantas vezes o padrão "100" aparece na sequência.

Restrições
1 ≤ N ≤ 10^4
1≤ N ≤ 10^4 */
package main
import f "fmt"
func main(){
var N int
f.Print("Digite o tamanho da sequência: ")
f.Scanf("%d", &N)
var seq string
f.Print("Digite a sequência de zeros e uns: ")
f.Scanf("%s", &seq)
var count int
for i := 0; i <= N-3; i++ {
if seq[i:i+3] == "100" {
count++
}
}f.Printf("O padrão '100' aparece %d vezes na sequência.\n", count)
}