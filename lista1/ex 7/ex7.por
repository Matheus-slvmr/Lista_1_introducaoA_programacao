programa
{
	inclua biblioteca Matematica --> mat
	
	funcao inicio()
	{
		
		real fahrenheit, polegadas
		real celsius, milimetros
		
		leia(fahrenheit)
		leia(polegadas)

		celsius = (5.0 * fahrenheit - 160.0) / 9.0
		milimetros = polegadas * 25.4
		escreva("O VALOR EM CELSIUS = ", mat.arredondar(celsius, 2), "\n")
		escreva("A QUANTIDADE DE CHUVA E = ", mat.arredondar(milimetros, 2), "\n")
	}
}
/* $$$ Portugol Studio $$$ 
 * 
 * Esta seção do arquivo guarda informações do Portugol Studio.
 * Você pode apagá-la se estiver utilizando outro editor.
 * 
 * @POSICAO-CURSOR = 393; 
 * @PONTOS-DE-PARADA = ;
 * @SIMBOLOS-INSPECIONADOS = ;
 * @FILTRO-ARVORE-TIPOS-DE-DADO = inteiro, real, logico, cadeia, caracter, vazio;
 * @FILTRO-ARVORE-TIPOS-DE-SIMBOLO = variavel, vetor, matriz, funcao;
 */