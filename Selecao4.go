package main

import "fmt"

func main() {

	var altura, peso float64
	var sexo int

	fmt.Print("Digite sua altura:")
	fmt.Scan(&altura)

	fmt.Print("Digite o peso:")
	fmt.Scan(&peso)

	fmt.Println("Se for homem digite 1\nSe for mulher digite 2")
	fmt.Print("Qual o seu sexo?")
	fmt.Scan(&sexo)
	imc := peso / (altura * altura)
	fmt.Println("O seu imc é de: ", imc)

	if sexo == 1 {
		if imc < 20 {
			fmt.Print("Você esta abaixo do normal")
		} else if imc < 25 {
			fmt.Print("Você esta normal")
		} else if imc > 25 {
			fmt.Print("Você esta obeso")
		}
	} else if sexo == 2 {
		if imc < 19 {
			fmt.Print("Você esta abaixo do normal")
		} else if imc < 24 {
			fmt.Print("Você esta normal")
		} else if imc > 24 {
			fmt.Print("Você esta obesa")
		}
	}
}
