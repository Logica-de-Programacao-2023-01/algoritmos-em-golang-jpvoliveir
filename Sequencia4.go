package main

import "fmt"

func main() {

	var num1, num2, num3 float64

	fmt.Println("Leia os seguintes numeros")
	fmt.Scan(&num1, &num2, &num3)
	fmt.Println("Qual a média ponderada entre eles?")
	fmt.Println(((num1*2 + num2*3 + num3*4) / (2 + 3 + 4)))

}
