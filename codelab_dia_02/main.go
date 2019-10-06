package main

import "fmt"

func main() {

	println("Exercicio 1")

	x := 3
	y := 5

	fmt.Println("O maior eh: ")
	if x > y {
		fmt.Printf("y = %v", x)
	} else if x < y {
		fmt.Printf("y = %v", y)
	} else {
		fmt.Print("sao iguais")
	}

	fmt.Println("Exercicio 2")
	ano := 2002

	if ano > (2019 - 16) {
		fmt.Println("Nao esta apto")
	} else {
		fmt.Println("Esta apto")
	}

}