package main

import "fmt"

func main() {

	var num1, num2 int
	fmt.Println("Leia o primmeiro numero:")
	fmt.Scan(&num1)
	fmt.Println("Leia o segundo numero:")
	fmt.Scan(&num2)

	if num1 >= 0 && num2 >= 0 {
		fmt.Print("O resultado da multiplição:", num1*num2)
	} else {
		fmt.Println("Resultado da soma", num1+num2)
	}
}
