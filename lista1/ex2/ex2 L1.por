programa
{
	inclua biblioteca Matematica
	
	funcao inicio()
	{
	inteiro numCasoTeste
	inteiro nv = 0
	escreva("Qual a quantidade de casos testes?")
	leia(numCasoTeste)

	enquanto(nv < numCasoTeste){
		
		nv = nv++
		real total
		real popular
		real geral 
		real arquibancadas
		real cadeiras
		escreva("Qual a quantiade de  ingressoa do jogo Nº ",nv)
		leia(total,popular,geral,arquibancadas,cadeiras)
		real calculoPopular = (popular/100)*(total)*(1)
		real calculoGeral = (geral/100)*(total)*(5)
		real calculoArquibancadas = (arquibancadas/100)*(total)*(10)
		real calculoCadeiras = (cadeiras/100)*(total)*(20)
		real soma = (calculoPopular + calculoGeral + calculoArquibancadas + calculoCadeiras)
	     escreva("A renda do jogo Nº", nv , " é :",soma,"\n")



		
		}	
	}
}

/* $$$ Portugol Studio $$$ 
 * 
 * Esta seção do arquivo guarda informações do Portugol Studio.
 * Você pode apagá-la se estiver utilizando outro editor.
 * 
 * @POSICAO-CURSOR = 777; 
 * @PONTOS-DE-PARADA = ;
 * @SIMBOLOS-INSPECIONADOS = ;
 * @FILTRO-ARVORE-TIPOS-DE-DADO = inteiro, real, logico, cadeia, caracter, vazio;
 * @FILTRO-ARVORE-TIPOS-DE-SIMBOLO = variavel, vetor, matriz, funcao;
 */