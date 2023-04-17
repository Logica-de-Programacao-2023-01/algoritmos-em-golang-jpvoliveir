package main

import "fmt"

func main() {

	var num1 int
	fmt.Print("digite um numero de 1 a 7:")
	fmt.Scan(&num1)

	switch num1 {
	case 1:
		fmt.Println("Domingo")
	case 2:
		fmt.Println("Segunda")
	case 3:
		fmt.Println("Terça")
	case 4:
		fmt.Println("Quarta")
	case 5:
		fmt.Println("Quinta")
	case 6:
		fmt.Println("Sexta")
	case 7:
		fmt.Println("Sabado")
	}

}
