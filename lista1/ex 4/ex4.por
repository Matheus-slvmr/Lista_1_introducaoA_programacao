programa
{
	inclua biblioteca Matematica
	
	funcao inicio()
	{
		escreva("Qual o salario minimo?\n")
		real salario
		real energia
		leia(salario)
		escreva("Quanto gasta de energia em sua residencia?\n")
		leia(energia)
		real calculoPorKW = (0.7 * salario)/ 100
		real calculoConsumo =  energia  * calculoPorKW 
		real valorDesconto
		valorDesconto = calculoConsumo * 0.90
		escreva("Custo por kW: R$"calculoPorKW, "\n")
		escreva("Custo do consumo: R$"calculoConsumo, "\n")
		escreva("Custo com desconto: R$"valorDesconto, "\n")
		
		
		
		
	}
}
/* $$$ Portugol Studio $$$ 
 * 
 * Esta seção do arquivo guarda informações do Portugol Studio.
 * Você pode apagá-la se estiver utilizando outro editor.
 * 
 * @POSICAO-CURSOR = 510; 
 * @PONTOS-DE-PARADA = ;
 * @SIMBOLOS-INSPECIONADOS = ;
 * @FILTRO-ARVORE-TIPOS-DE-DADO = inteiro, real, logico, cadeia, caracter, vazio;
 * @FILTRO-ARVORE-TIPOS-DE-SIMBOLO = variavel, vetor, matriz, funcao;
 */