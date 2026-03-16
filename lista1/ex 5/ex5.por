programa
{
	funcao inicio()
	{
		inteiro conta
		real consumo, valor_conta = 0.0
		caracter tipo

		leia(conta, consumo, tipo)

		se (tipo == 'R' ou tipo == 'r') 
		{
			
			valor_conta = 5.00 + (consumo * 0.05)
		}
		senao se (tipo == 'C' ou tipo == 'c')
		{
			
			valor_conta = 500.00 + (consumo * 0.25)
		}
		senao se (tipo == 'I' ou tipo == 'i')
		{
			
			valor_conta = 800.00 + (consumo * 0.04)
		}

		escreva("CONTA = ", conta, "\n")		
		escreva("VALOR DA CONTA = ", valor_conta, "\n")
	}
}
/* $$$ Portugol Studio $$$ 
 * 
 * Esta seção do arquivo guarda informações do Portugol Studio.
 * Você pode apagá-la se estiver utilizando outro editor.
 * 
 * @POSICAO-CURSOR = 493; 
 * @PONTOS-DE-PARADA = ;
 * @SIMBOLOS-INSPECIONADOS = ;
 * @FILTRO-ARVORE-TIPOS-DE-DADO = inteiro, real, logico, cadeia, caracter, vazio;
 * @FILTRO-ARVORE-TIPOS-DE-SIMBOLO = variavel, vetor, matriz, funcao;
 */