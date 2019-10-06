package main

import "fmt"

func main() {

	println("Exercício 1 ")
	exercicio1()

}

func exercicio1() {
	/*
		Utilizando a estrutura for com regra de incremento +1,
		faça um programa que printe na tela os números inteiros
		de 1 a 100 que sejam múltiplos de 3.
	 */
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 {
			fmt.Println(i)
		}
	}
}