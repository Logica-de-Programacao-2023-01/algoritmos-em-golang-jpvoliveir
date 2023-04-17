package main

import "fmt"

func main() {

	var num1 int
	fmt.Print("Digite um numero inteiro:")
	fmt.Scan(&num1)

	if num1%3 == 0 && num1%5 == 0 {
		fmt.Print("o numero é multiplo de 3 e de 5 ao mesmo tempo")
	} else {
		fmt.Print("o numero não é  multiplo de 3 e de 5")
	}
}
