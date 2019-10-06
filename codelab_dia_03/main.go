package main

import "fmt"

func main() {

	println("Exercício 1 ")
	exercicio1()
	println("Exercício 2 ")
	exercicio2()

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

func exercicio2()  {
	/*
		Escreva um programa que contenha uma função que
		teste se um dado número inteiro é múltiplo de 5 ou não.
	 */
	numero := 55
	if numero % 5 == 0 {
		fmt.Printf("O numero %v é multiplo de 5", numero)
	} else {
		fmt.Printf("O numero %v não é multiplo de 5", numero)
	}
}