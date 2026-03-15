programa
{
	
	inclua biblioteca Matematica --> mat
	
	funcao inicio()
	{
	
		real a, b, c, d, determinante


		leia(a)
		leia(b)
		leia(c)
		leia(d)

		
		determinante = (a * d) - (b * c)

		
		escreva("O VALOR DO DETERMINANTE E = ", mat.arredondar(determinante, 2), "\n")
	}
}
/* $$$ Portugol Studio $$$ 
 * 
 * Esta seção do arquivo guarda informações do Portugol Studio.
 * Você pode apagá-la se estiver utilizando outro editor.
 * 
 * @POSICAO-CURSOR = 152; 
 * @PONTOS-DE-PARADA = ;
 * @SIMBOLOS-INSPECIONADOS = ;
 * @FILTRO-ARVORE-TIPOS-DE-DADO = inteiro, real, logico, cadeia, caracter, vazio;
 * @FILTRO-ARVORE-TIPOS-DE-SIMBOLO = variavel, vetor, matriz, funcao;
 */