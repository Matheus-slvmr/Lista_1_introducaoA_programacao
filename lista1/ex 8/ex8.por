programa
{
	inclua biblioteca Matematica --> mat
	
	funcao inicio()
	{
		
		real raio, altura, custo_por_m2
		real area_circulo, area_lateral, area_total, custo_final
		const real PI = 3.14159

		custo_por_m2 = 100.0

	
		leia(raio)
		leia(altura)

		// 2. Cálculos Geométricos
		// Área do círculo (tampa): Ac = π * r²
		area_circulo = PI * (raio * raio)

		// Área lateral: Al = 2 * π * r * a
		area_lateral = 2 * PI * raio * altura

		// Área total: At = 2 * Ac + Al
		area_total = (2 * area_circulo) + area_lateral

		
		custo_final = area_total * custo_por_m2

		
		escreva("O VALOR DO CUSTO E = ", mat.arredondar(custo_final, 2), "\n")
	}
}
/* $$$ Portugol Studio $$$ 
 * 
 * Esta seção do arquivo guarda informações do Portugol Studio.
 * Você pode apagá-la se estiver utilizando outro editor.
 * 
 * @POSICAO-CURSOR = 568; 
 * @PONTOS-DE-PARADA = ;
 * @SIMBOLOS-INSPECIONADOS = ;
 * @FILTRO-ARVORE-TIPOS-DE-DADO = inteiro, real, logico, cadeia, caracter, vazio;
 * @FILTRO-ARVORE-TIPOS-DE-SIMBOLO = variavel, vetor, matriz, funcao;
 */