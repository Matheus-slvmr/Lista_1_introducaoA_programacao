programa
{
	inclua biblioteca Matematica -->mat
	
	funcao inicio()
	{
		real soma = 0.0
		inteiro n
		leia (n)
		
		se (n>1)
		{
				para(inteiro k = 1; k <= n;k++){
					
					soma = soma + (1.0 / k)
					
					}
					escreva (mat.arredondar(soma,6),"\n")
			}senao
			{
				escreva("invalido\n")
			}


		
	}


	
}
/* $$$ Portugol Studio $$$ 
 * 
 * Esta seção do arquivo guarda informações do Portugol Studio.
 * Você pode apagá-la se estiver utilizando outro editor.
 * 
 * @POSICAO-CURSOR = 305; 
 * @PONTOS-DE-PARADA = ;
 * @SIMBOLOS-INSPECIONADOS = ;
 * @FILTRO-ARVORE-TIPOS-DE-DADO = inteiro, real, logico, cadeia, caracter, vazio;
 * @FILTRO-ARVORE-TIPOS-DE-SIMBOLO = variavel, vetor, matriz, funcao;
 */