package main

import "fmt"

func main() {

	var num1 float64
	fmt.Print("Quantos anos você tem?")
	fmt.Scan(&num1)
	fmt.Print("você tem", (num1 * 365), "dias de vida")
}
