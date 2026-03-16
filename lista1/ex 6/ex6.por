programa
{
	inclua biblioteca Matematica --> mat
	
	funcao inicio()
	{	
		inteiro n
		real fahrenheit, celsius
		leia(n)

		para (inteiro i = 0; i < n; i++)
		{
			leia(fahrenheit)

			
			celsius = (5.0 * (fahrenheit - 32.0)) / 9.0
			escreva(mat.arredondar(fahrenheit, 2), " FAHRENHEIT EQUIVALE A ", mat.arredondar(celsius, 2), " CELSIUS\n")
		}
	}
}
/* $$$ Portugol Studio $$$ 
 * 
 * Esta seção do arquivo guarda informações do Portugol Studio.
 * Você pode apagá-la se estiver utilizando outro editor.
 * 
 * @POSICAO-CURSOR = 352; 
 * @PONTOS-DE-PARADA = ;
 * @SIMBOLOS-INSPECIONADOS = ;
 * @FILTRO-ARVORE-TIPOS-DE-DADO = inteiro, real, logico, cadeia, caracter, vazio;
 * @FILTRO-ARVORE-TIPOS-DE-SIMBOLO = variavel, vetor, matriz, funcao;
 */