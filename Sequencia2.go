package main

import "fmt"

func main() {
	var num1 int

	fmt.Println("qual é o número?")
	fmt.Scan(&num1)

	fmt.Print("o dobro do seu número é", num1*2, "o triplo do seu numero é", num1*3, "o quadruplo do seu numero é", num1*4)
}
