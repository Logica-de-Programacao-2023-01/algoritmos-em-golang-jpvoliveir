package main

import "fmt"

func main() {
	var num1, num2, num3 int
	fmt.Println("leia os seguintes números:")
	fmt.Scan(&num1, &num2, &num3)

	if num1 < num2 && num1 < num3 {
		fmt.Printf("o numero %d é menor", num1)
	} else if num2 < num1 && num2 < num3 {
		fmt.Printf("o numero %d é menor", num2)

	} else {
		fmt.Printf("O numero %d é menor", num3)
	}
}
