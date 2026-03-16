programa
{
	
	funcao inicio()
	{
	inteiro x, y , z, conversaoHS, conversaoMS, segundos
		escreva("a hora em hora(x) minuto(y) segundo(z): ")
		leia(x , y , z)
		conversaoHS = x *( 60 * 60)
		conversaoMS = y * 60
		segundos = conversaoHS + conversaoMS + z 
		escreva("O TEMPO EM SEGUNDOS E =" , segundos,"\n")
	}
}
/* $$$ Portugol Studio $$$ 
 * 
 * Esta seção do arquivo guarda informações do Portugol Studio.
 * Você pode apagá-la se estiver utilizando outro editor.
 * 
 * @POSICAO-CURSOR = 321; 
 * @PONTOS-DE-PARADA = ;
 * @SIMBOLOS-INSPECIONADOS = ;
 * @FILTRO-ARVORE-TIPOS-DE-DADO = inteiro, real, logico, cadeia, caracter, vazio;
 * @FILTRO-ARVORE-TIPOS-DE-SIMBOLO = variavel, vetor, matriz, funcao;
 */