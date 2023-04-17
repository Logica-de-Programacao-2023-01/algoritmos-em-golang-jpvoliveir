package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Print("leia os seguintes numeros:")
	fmt.Scan(&num1, &num2)

	if num1 > num2 {
		fmt.Printf("o numero %d é maior", num1)
	} else if num2 > num1 {
		fmt.Printf("o numero %d é maior", num2)

	} else {
		fmt.Print("Os numeros são iguais")
	}
}
