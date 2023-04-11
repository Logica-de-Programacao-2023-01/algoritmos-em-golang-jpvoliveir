package main

import "fmt"

func main() {
	var num1 float64

	fmt.Println("Digite peso:")
	fmt.Scan(&num1)
	fmt.Print("O seu peso em libras é=", num1*2.20462)

}
