package main

import (
	"fmt"
)

func encontrarDivisores(numero int) []int {
	divisores := []int{}

	for i := 1; i <= numero; i++ {
		if numero%i == 0 {
			divisores = append(divisores, i)
		}
	}

	return divisores
}

func main() {
	var numero int
	fmt.Print("Digite um número inteiro positivo: ")
	fmt.Scanln(&numero)

	if numero <= 0 {
		fmt.Println("O número deve ser inteiro positivo.")
		return
	}

	divisores := encontrarDivisores(numero)

	fmt.Printf("Divisores de %d: %v\n", numero, divisores)
}
