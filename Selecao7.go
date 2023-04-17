package main

import "fmt"

func main() {

	var num1 float64
	fmt.Println("Digite o salário:")
	fmt.Scan(&num1)

	var num2 float64
	if num1 < 1000.0 {
		num2 = 0.1
	} else {
		num2 = 0.05
	}
	novoSalario := num1 * (1 + num2)
	fmt.Printf("O novo salário do funcionário é de R$ %.2f", novoSalario)
}
