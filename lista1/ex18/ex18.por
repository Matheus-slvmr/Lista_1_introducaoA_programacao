programa
{
	funcao inicio()
	{
		inteiro a1, r, n
		
		inteiro soma = 0
		inteiro termo
		
		leia(a1, r, n)
		termo = a1

		para (inteiro i = 0; i < n; i++)
		{
			soma = soma + termo
			termo = termo + r
		}

		escreva(soma, "\n")
	}
}
/* $$$ Portugol Studio $$$ 
 * 
 * Esta seção do arquivo guarda informações do Portugol Studio.
 * Você pode apagá-la se estiver utilizando outro editor.
 * 
 * @POSICAO-CURSOR = 183; 
 * @PONTOS-DE-PARADA = ;
 * @SIMBOLOS-INSPECIONADOS = ;
 * @FILTRO-ARVORE-TIPOS-DE-DADO = inteiro, real, logico, cadeia, caracter, vazio;
 * @FILTRO-ARVORE-TIPOS-DE-SIMBOLO = variavel, vetor, matriz, funcao;
 */