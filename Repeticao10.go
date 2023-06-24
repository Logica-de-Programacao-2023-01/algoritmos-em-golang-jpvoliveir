package main

import (
	"fmt"
)

func main() {
	var (
		numero   int
		maior    int
		primeiro = true
	)

	for {
		fmt.Print("Digite um número inteiro (0 para sair): ")
		fmt.Scanln(&numero)

		if numero == 0 {
			break
		}

		if primeiro || numero > maior {
			maior = numero
			primeiro = false
		}
	}

	if primeiro {
		fmt.Println("Nenhum número foi digitado.")
	} else {
		fmt.Println("O maior número digitado é:", maior)
	}
}
