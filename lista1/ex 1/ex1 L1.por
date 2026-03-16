programa
{
	inclua biblioteca Matematica
	
	funcao inicio()
	{
	
		real N1,N2,N3
		
		escreva("Digite sua primeira nota:")
		leia(N1)
		escreva("Digite sua segunda nota:")
		leia(N2)
		escreva("Digite sua terceira nota:")
		leia(N3)
		real M = (N1+N2+N3)/3
		
		se (M >= 6.0){
			escreva("\nSua media é ",M)
			escreva("\nAprovado")
			}senao{
				escreva("\nSua media é",M)
				escreva("\nReprovado")
				}
		
		



		
	}
}
/* $$$ Portugol Studio $$$ 
 * 
 * Esta seção do arquivo guarda informações do Portugol Studio.
 * Você pode apagá-la se estiver utilizando outro editor.
 * 
 * @POSICAO-CURSOR = 394; 
 * @PONTOS-DE-PARADA = ;
 * @SIMBOLOS-INSPECIONADOS = ;
 * @FILTRO-ARVORE-TIPOS-DE-DADO = inteiro, real, logico, cadeia, caracter, vazio;
 * @FILTRO-ARVORE-TIPOS-DE-SIMBOLO = variavel, vetor, matriz, funcao;
 */