programa {
	funcao inicio() {
		inteiro a, b, c, aux

		escreva("Digite três valores:\n")
		leia(a, b, c)

		
		se (a > b) { aux = a; a = b; b = aux }
		se (a > c) { aux = a; a = c; c = aux }
		se (b > c) { aux = b; b = c; c = aux }

		escreva("Valores ordenados: ", a, ", ", b, ", ", c)
	}
}
/* $$$ Portugol Studio $$$ 
 * 
 * Esta seção do arquivo guarda informações do Portugol Studio.
 * Você pode apagá-la se estiver utilizando outro editor.
 * 
 * @POSICAO-CURSOR = 149; 
 * @PONTOS-DE-PARADA = ;
 * @SIMBOLOS-INSPECIONADOS = ;
 * @FILTRO-ARVORE-TIPOS-DE-DADO = inteiro, real, logico, cadeia, caracter, vazio;
 * @FILTRO-ARVORE-TIPOS-DE-SIMBOLO = variavel, vetor, matriz, funcao;
 */