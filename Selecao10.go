package main

import "fmt"

func main() {

	var idade int
	fmt.Println("Quantos anos voce tem?")
	fmt.Scan(&idade)

	if idade <= 9 {
		fmt.Print("Mirim")
	} else if idade <= 13 {
		fmt.Print("Infantil")
	} else if idade <= 17 {
		fmt.Print("Juvenil")
	} else {
		fmt.Print("Adulto")
	}

}
