programa
{
	funcao inicio()
	{
		
		inteiro n1, n2, n3, numeroFinal, quadrado

		leia(n1)
		leia(n2)
		leia(n3)

		
		se (n1 < 0 ou n1 > 9 ou n2 < 0 ou n2 > 9 ou n3 < 0 ou n3 > 9)// se nx for menor que zero ou maior que nove é invalido senao vai pro calculo
		{
			escreva("DIGITO INVALIDO\n")
		}
		senao
		{
			
			numeroFinal = (n1 * 100) + (n2 * 10) + n3
			quadrado = numeroFinal * numeroFinal
			escreva(numeroFinal, ", ", quadrado, "\n")
		}
	}
}
/* $$$ Portugol Studio $$$ 
 * 
 * Esta seção do arquivo guarda informações do Portugol Studio.
 * Você pode apagá-la se estiver utilizando outro editor.
 * 
 * @POSICAO-CURSOR = 453; 
 * @PONTOS-DE-PARADA = ;
 * @SIMBOLOS-INSPECIONADOS = ;
 * @FILTRO-ARVORE-TIPOS-DE-DADO = inteiro, real, logico, cadeia, caracter, vazio;
 * @FILTRO-ARVORE-TIPOS-DE-SIMBOLO = variavel, vetor, matriz, funcao;
 */