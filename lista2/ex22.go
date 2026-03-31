/* Desenvolver um programa que calcule o salário bruto e o salário líquido de um funcionário.
• Dados de Entrada: Matrícula do funcionário (int);
Quantidade de horas-extras trabalhadas.
• Constantes: Salário Mínimo = R$ 788.00;
Valor da Hora-Extra = R$ 10.00.
Sabe-se:
• Salário hora-extra = horas-extras * Valor da Hora-Extra;
• Salário bruto = 3 * Salário Mínimo + Salário hora-extra;
• Desconto INSS = 12 % do salário bruto, se salário bruto for maior que R$ 1500,00;
• Desconto do Imposto de Renda = 20 % do Salário Bruto, se o mesmo for maior que R$ 2000,00;
• Salário líquido = salário bruto – deduções. */
package main
import (
	"fmt"
)
func main() {
	const salario_minimo float64 = 788.00
	const valor_hora_extra float64 = 10.00
	var matricula int
	fmt.Print("Digite a matrícula do funcionário: ")
	fmt.Scan(&matricula)
	var horas_extras int
	fmt.Print("Digite a quantidade de horas-extras trabalhadas: ")
	fmt.Scan(&horas_extras)
	salario_hora_extra := float64(horas_extras) * valor_hora_extra
	salario_bruto := 3*salario_minimo + salario_hora_extra
	var desconto_inss float64
	if salario_bruto > 1500.00 {
		desconto_inss = salario_bruto * 0.12
	} else {
		desconto_inss = 0
	}
	var desconto_ir float64
	if salario_bruto > 2000.00 {
		desconto_ir = salario_bruto * 0.20
	} else {
		desconto_ir = 0
	}
	salario_liquido := salario_bruto - desconto_inss - desconto_ir
	fmt.Printf("Matrícula do funcionário: %d\n", matricula)
	fmt.Printf("Salário bruto: %.2f\n", salario_bruto)
	fmt.Printf("Desconto INSS: %.2f\n", desconto_inss)
	fmt.Printf("Desconto Imposto de Renda: %.2f\n", desconto_ir)
	fmt.Printf("Salário líquido: %.2f\n", salario_liquido)
}