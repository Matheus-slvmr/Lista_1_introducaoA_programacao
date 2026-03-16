programa
{
	
	funcao inicio()
	{
		inteiro n
		leia(n)
		se (n> 5 e n<2000){
			para (inteiro i = 2; i <= n; i = i + 2)
			{
				escreva(i, "^2 = ", i * i, "\n")
			}
		}senao{
			escreva("Numero invalido!\n")
			 }
	}
}
/* $$$ Portugol Studio $$$ 
 * 
 * Esta seção do arquivo guarda informações do Portugol Studio.
 * Você pode apagá-la se estiver utilizando outro editor.
 * 
 * @POSICAO-CURSOR = 224; 
 * @PONTOS-DE-PARADA = ;
 * @SIMBOLOS-INSPECIONADOS = ;
 * @FILTRO-ARVORE-TIPOS-DE-DADO = inteiro, real, logico, cadeia, caracter, vazio;
 * @FILTRO-ARVORE-TIPOS-DE-SIMBOLO = variavel, vetor, matriz, funcao;
 */