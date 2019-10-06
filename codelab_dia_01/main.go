package main

import (
	"fmt"
)

func main() {

	// exercício 1
	var numero_um int
	numero_um = 1
	fmt.Println(numero_um)

	// exercício 2
	A := 230
	B := 27
	fmt.Println("A: ", A)
	fmt.Println("B: ", B)

	soma := A + B
	subtracao := A - B
	fmt.Println("Soma: ", soma)
	fmt.Println("Subtração: ", subtracao)

	// exercício 3
	multiplicacao := A * B
	divisao := A / B
	fmt.Println("Multiplicação: ", multiplicacao)
	fmt.Println("Divisão: ", divisao)
	fmt.Println("_______________________")

	// exercício 4
	fmt.Printf("Soma: %v \nSubtração: %v \nMultiplicação: %v \nDivisão: %v",
		soma, subtracao, multiplicacao, divisao)
	fmt.Println("_______________________")

	// exercicio 5

	//Preços dos itens do mercado
	var preçoDoPão float64 = 4.60
	//var preçoDaAveia float64 = 5
	var preçoDoAzeite float64 = 19.95
	var preçoDoSuco float64 = 7
	var preçoDaÁgua float64 = 2.15
	var preçoDoKGMaçã float64 = 8.90
	//var preçoDoKGBanana float64 = 4.5

	totalCompras := 3 * preçoDoPão
	totalCompras += preçoDoAzeite
	totalCompras += 2 * preçoDoSuco
	totalCompras += 5 * preçoDaÁgua
	totalCompras += 1.5 * preçoDoKGMaçã

	fmt.Println("Total das compras: ", totalCompras)
	fmt.Println("_______________________")

	// Exercicio 6
	// nao eh possivel utilizar palavras reservadas

	// Exercicio 7
	var X int = 15
	Y := 31
	var Z int
	Z = 47

	fmt.Println("\n\t X \t Y \t Z")
	fmt.Printf("base 2 : %b \t%b \t%b", X, Y, Z)
	fmt.Printf("\nbase 10: %d \t%d \t%d", X, Y, Z)
	fmt.Printf("\nbase 16: %#X \t%#X \t%#X\n", X, Y, Z)
	fmt.Println("_______________________")

	// Exercicio 7
	// Nao 'e possivel compilar o programa com variaveis nao utilizadas

}