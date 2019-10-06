package main

import "fmt"

func main() {

	println("Exercicio 1")

	x := 3
	y := 5

	println("O maior eh: ")
	if x > y {
		fmt.Printf("y = %v", x)
	} else if x < y {
		fmt.Printf("y = %v", y)
	} else {
		fmt.Print("sao iguais")
	}

}