package main

import "fmt"

func main() {
	var num1 float64

	fmt.Println("Preço do produto:")
	fmt.Scan(&num1)
	fmt.Print("Com 10% de desconto o salario passa a ser:", num1-(num1*0.10), "R$")

}
