programa
{
	
	funcao inicio()
	{
	real salario
	real reajuste
		escreva("Seu salario:")
		leia(salario)
		se (salario <= 300){
			escreva("seu reajuste:", salario + (0.5 * salario))
			}senao se (salario >= 500){
				escreva("seu reajuste:", salario + (0.3 * salario))
				}
	}
}
/* $$$ Portugol Studio $$$ 
 * 
 * Esta seção do arquivo guarda informações do Portugol Studio.
 * Você pode apagá-la se estiver utilizando outro editor.
 * 
 * @POSICAO-CURSOR = 286; 
 * @PONTOS-DE-PARADA = 11;
 * @SIMBOLOS-INSPECIONADOS = ;
 * @FILTRO-ARVORE-TIPOS-DE-DADO = inteiro, real, logico, cadeia, caracter, vazio;
 * @FILTRO-ARVORE-TIPOS-DE-SIMBOLO = variavel, vetor, matriz, funcao;
 */