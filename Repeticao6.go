package main

import "fmt"

func main() {
	var num int
	fmt.Println("Digite um numero:")
	fmt.Scan(&num)

	for i := 0; i <= 10; i++ {
		fmt.Println(num * i)
	}

}
