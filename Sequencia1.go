package main

import "fmt"

func main() {
	var num1, num2, num3 int

	fmt.Println("qual é o primeiro número?")
	fmt.Scan(&num1)
	fmt.Println("qual é o segundo número?")
	fmt.Scan(&num2)
	fmt.Println("qual é o terceiro número?")
	fmt.Scan(&num3)

	fmt.Print("A soma dos seus 3 números é: ", num1+num2+num3)
}
