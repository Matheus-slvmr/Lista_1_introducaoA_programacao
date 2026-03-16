/*
 É entre 9 e 10? Atribua 'A'.

É entre 7.5 e 8.9? Atribua 'B'.

É entre 6 e 7.4? Atribua 'C'.

É menor que 6? Atribua 'D'.
 */
programa
{
	
	funcao inicio()
	{
	real nota
		escreva("qual sua nota??")
		leia (nota)
		se (nota >= 9) {
			escreva("seu conceito é A")
			}senao se (nota >= 7.5){
				escreva("seu conceito é B")
				}senao se (nota >= 6) {
					escreva("seu conceito é C")
					}senao {
						escreva("seu conceito é D")
						}
	}
}
/* $$$ Portugol Studio $$$ 
 * 
 * Esta seção do arquivo guarda informações do Portugol Studio.
 * Você pode apagá-la se estiver utilizando outro editor.
 * 
 * @POSICAO-CURSOR = 406; 
 * @PONTOS-DE-PARADA = ;
 * @SIMBOLOS-INSPECIONADOS = ;
 * @FILTRO-ARVORE-TIPOS-DE-DADO = inteiro, real, logico, cadeia, caracter, vazio;
 * @FILTRO-ARVORE-TIPOS-DE-SIMBOLO = variavel, vetor, matriz, funcao;
 */