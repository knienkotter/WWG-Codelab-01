package main

import "fmt"

func main() {

	println("Exercício 1 ")
	exercicio1()
	println("Exercício 2 ")
	exercicio2()
	println("Exercício 3 ")
	exercicio3()
	println("Exercício 4 ")
	exercicio4(3, 7, 92)
	println("Exercício 5 ")
	exercicio5(4, 26)
	exercicio5(3, 7, 92)

}

func exercicio1() {
	/*
		Utilizando a estrutura for com regra de incremento +1,
		faça um programa que printe na tela os números inteiros
		de 1 a 100 que sejam múltiplos de 3.
	*/
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

func exercicio2() {
	/*
		Repita o que foi solicitado no exercício 01, mas dessa
		vez **printe na tela apenas os múltiplos de 9** e utilize
		a keyword **continue** na sua condicional.
	*/
	for i := 1; i <= 100; i++ {
		if i%9 != 0 {
			continue
		}
		fmt.Println(i)
	}
}

func exercicio3() {
	/*
		Escreva um programa que contenha uma função que
		teste se um dado número inteiro é múltiplo de 5 ou não.
	*/
	numero := 55
	if numero%5 == 0 {
		fmt.Printf("O numero %v é multiplo de 5", numero)
	} else {
		fmt.Printf("O numero %v não é multiplo de 5", numero)
	}
	fmt.Println()
}

func exercicio4(a int, b int, c int) {
	/*
		1) Escreva uma função que:
			1) Receba como parâmetro três números inteiros, do tipo _int_;
			2) Some-os; e
			3) Printe na tela o resultado da soma.
		2) Faça uma chamada dessa função na _func main()_.
	*/
	soma := a + b + c
	fmt.Println("Resultado da soma: ", soma)
}

func exercicio5(v ...int) {
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
