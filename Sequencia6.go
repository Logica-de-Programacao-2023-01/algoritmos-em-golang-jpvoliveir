package main

import "fmt"

func main() {

	var num1 float64
	fmt.Println("salario inicial:")
	fmt.Scan(&num1)
	fmt.Print("o seu novo salário é=", (num1*0.15)+num1)
}
