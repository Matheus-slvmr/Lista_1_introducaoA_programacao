programa
{
	funcao inicio()
	{
		inteiro x, y
		
		leia(x, y)

		
		se (x % 2 != 0) 
		{
			escreva("O PRIMEIRO NUMERO NAO E PAR\n")
		}
		senao
		{
			
			para (inteiro i = 0; i < y; i++)
			{
				escreva(x, " ")
				x = x + 2 
			}
			escreva("\n")
		}
	}
}
/* $$$ Portugol Studio $$$ 
 * 
 * Esta seção do arquivo guarda informações do Portugol Studio.
 * Você pode apagá-la se estiver utilizando outro editor.
 * 
 * @POSICAO-CURSOR = 48; 
 * @PONTOS-DE-PARADA = ;
 * @SIMBOLOS-INSPECIONADOS = ;
 * @FILTRO-ARVORE-TIPOS-DE-DADO = inteiro, real, logico, cadeia, caracter, vazio;
 * @FILTRO-ARVORE-TIPOS-DE-SIMBOLO = variavel, vetor, matriz, funcao;
 */