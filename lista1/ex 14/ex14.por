programa
{
	inclua biblioteca Matematica --> mat
	
	funcao inicio()
	{
	real alturaPiramide
	real arestaHexa
	real volumePiramide
	real areaHexa
		leia(alturaPiramide, arestaHexa)
		areaHexa = (3 * (arestaHexa * arestaHexa) * mat.raiz(3.0, 2.0)) / 2
		volumePiramide = (areaHexa * alturaPiramide) / 3
		escreva("O VOLUME DA PIRAMIDE E = ", mat.arredondar(volumePiramide, 2), " METROS CUBICOS\n")
		
	}
}
/* $$$ Portugol Studio $$$ 
 * 
 * Esta seção do arquivo guarda informações do Portugol Studio.
 * Você pode apagá-la se estiver utilizando outro editor.
 * 
 * @POSICAO-CURSOR = 340; 
 * @PONTOS-DE-PARADA = ;
 * @SIMBOLOS-INSPECIONADOS = ;
 * @FILTRO-ARVORE-TIPOS-DE-DADO = inteiro, real, logico, cadeia, caracter, vazio;
 * @FILTRO-ARVORE-TIPOS-DE-SIMBOLO = variavel, vetor, matriz, funcao;
 */