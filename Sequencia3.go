package main

import "fmt"

func main() {
	var altura float32
	var peso float32

	fmt.Println("qual a sua altura?")
	fmt.Scan(&altura)
	fmt.Println("qual o seu peso?")
	fmt.Scan(&peso)

	fmt.Print("o seu imc é:", peso/(altura*altura))

}
