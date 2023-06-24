package main

import "fmt"

func main() {
	i := 0
	for i <= 19 {
		if i%2 == 1 {
			fmt.Println(i)
		}
		i++
	}

}
