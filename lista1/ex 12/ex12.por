programa
{
	inclua biblioteca Matematica-->mat
	
	funcao inicio()
	{
	inteiro hTotais
	inteiro grupoDe3Horas
	inteiro horasSoltas
	real x
		escreva("quanto tempo usou a charrete em horas?")
		leia(hTotais)
		grupoDe3Horas = hTotais / 3
		horasSoltas = hTotais % 3
		x = (grupoDe3Horas * 10.0) + (horasSoltas * 5.0)
		escreva("o valor a pagar é:",mat.arredondar(x, 2))
		
	}
}
/* $$$ Portugol Studio $$$ 
 * 
 * Esta seção do arquivo guarda informações do Portugol Studio.
 * Você pode apagá-la se estiver utilizando outro editor.
 * 
 * @POSICAO-CURSOR = 281; 
 * @PONTOS-DE-PARADA = ;
 * @SIMBOLOS-INSPECIONADOS = ;
 * @FILTRO-ARVORE-TIPOS-DE-DADO = inteiro, real, logico, cadeia, caracter, vazio;
 * @FILTRO-ARVORE-TIPOS-DE-SIMBOLO = variavel, vetor, matriz, funcao;
 */