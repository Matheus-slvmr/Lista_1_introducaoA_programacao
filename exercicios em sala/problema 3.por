programa
{
	inclua biblioteca Matematica
 -- mat
	
	funcao inicio()
	{
	real v , d , calculo
	cadeia comando = ""
		enquanto (comando != "sair")
		{
			escreva("valor da mercadoria:")
		leia (v)
		escreva("valor do desconto em porcentagem:")
		leia(d)
		calculo = v - (v * (d / 100))
		escreva("com um desconto de ",d," % ","seu novo valor é R$ " ,calculo,"\n")
		
			}
	}
}

/* $$$ Portugol Studio $$$ 
 * 
 * Esta seção do arquivo guarda informações do Portugol Studio.
 * Você pode apagá-la se estiver utilizando outro editor.
 * 
 * @POSICAO-CURSOR = 375; 
 * @PONTOS-DE-PARADA = ;
 * @SIMBOLOS-INSPECIONADOS = ;
 * @FILTRO-ARVORE-TIPOS-DE-DADO = inteiro, real, logico, cadeia, caracter, vazio;
 * @FILTRO-ARVORE-TIPOS-DE-SIMBOLO = variavel, vetor, matriz, funcao;
 */