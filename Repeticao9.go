package main

import (
	"fmt"
)

func main() {
	var (
		numero   int
		soma     int
		contador int
	)

	for {
		fmt.Print("Digite um número inteiro (0 para sair): ")
		fmt.Scanln(&numero)

		if numero == 0 {
			break
		}

		soma += numero
		contador++
	}

	if contador == 0 {
		fmt.Println("Nenhum número foi digitado.")
	} else {
		media := float64(soma) / float64(contador)
		fmt.Printf("A média dos números digitados é: %.2f\n", media)
	}
}
