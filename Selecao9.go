package main

import "fmt"

func main() {
	var num1, num2, num3 int
	fmt.Println("Digite o primeiro numero:")
	fmt.Scan(&num1)
	fmt.Println("Digite o segundo numero:")
	fmt.Scan(&num2)
	fmt.Println("Digite o terceiro numero:")
	fmt.Scan(&num3)

	if num1 <= num2 && num1 <= num3 {
		if num2 <= num3 {
			fmt.Println(num1, num2, num3)
		} else {
			fmt.Println(num1, num3, num2)
		}
	} else if num2 <= num1 && num2 <= num3 {
		if num1 <= num3 {
			fmt.Println(num2, num1, num3)
		} else {
			fmt.Println(num2, num3, num1)
		}
	} else if num3 <= num1 && num3 <= num2 {
		if num1 <= num2 {
			fmt.Println(num3, num1, num2)
		} else {
			fmt.Println(num3, num2, num1)
		}
	}
}
