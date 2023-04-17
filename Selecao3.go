package main

import "fmt"

func main() {
	var num1 int
	fmt.Println("Leia o seguinte numero:")
	fmt.Scan(&num1)

	if num1%2 == 0 {
		fmt.Println("O número digitado é par.")
	} else {
		fmt.Println("O número digitado é ímpar.")
	}
}
