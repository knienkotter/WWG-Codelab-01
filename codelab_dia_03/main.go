package main

import "fmt"

func main() {

	println("Exercício 1 ")
	exercicio1()
	println("Exercício 2 ")
	exercicio2()
	println("Exercício 3 ")
	exercicio3(4, 6, 18)
	println("Exercício 4 ")
	exercicio4(4, 26)
	exercicio4(3, 7, 92)

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
	fmt.Println()
}

func exercicio3(a int, b int , c int)  {
	/*
		Escreva um programa que contenha uma função
		que some três operandos e que printe o resultado na tela.
	 */
	soma := a + b + c
	fmt.Println("Resultado da soma: ", soma)
}

func exercicio4(v ...int)  {
	/*
		Escreva um programa que contenha uma função que
		some n operandos e que printe o resultado na tela.
	 */
	fmt.Print(v, " ")
	total := 0
	for _, num := range v {
		total += num
	}
	fmt.Println(total)
}