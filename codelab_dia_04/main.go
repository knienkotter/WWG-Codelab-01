package main

import "fmt"

func main() {
	println("Exercício 1 ")
	lista := exercicio1()
	println("Exercício 2 ")
	lista_slice := lista[0:len(lista)] /* converte a lista em slice */
	exercicio2(lista_slice)
	println("Exercício 3 ")
	exercicio3()
	println("Exercício 4 ")
	exercicio4()
	println("Exercício 5 ")
	exercicio5()

}

func exercicio1() [10]int {

	var lista [10]int

	lista[0] = 12
	lista[1] = 23
	lista[2] = 34
	lista[3] = 45
	lista[4] = 56
	lista[5] = 67
	lista[6] = 78
	lista[7] = 89
	lista[8] = 90
	lista[9] = 01

	for indice, valor := range lista {
		fmt.Printf("índice: %d; valor: %d\n", indice, valor)
	}
	return lista
}

func exercicio2(lista []int) {

	var soma int
	for _, valor := range lista {
		soma = soma + valor
	}
	fmt.Println(soma)

}

func exercicio3() {

	lista := make([]string, 5)
	lista[0] = "Ismália"
	lista[1] = "Eminencia Parda"
	lista[2] = "AmarElo"
	lista[3] = "Libre"
	lista[4] = "Principia"

	for indice, item := range lista {
		fmt.Printf("índice %d) %s\n", indice, item)
	}

}

func exercicio4() {
	numeros := []int{7, 14, 21, 28, 35, 42, 49}
	for _, valor := range numeros {
		fmt.Printf("%d ", valor)
	}
	fmt.Println("")
	numeros = append(numeros, 56, 63, 70)
	for _, valor := range numeros {
		fmt.Printf("%d ", valor)
	}
	fmt.Println("")
}

func exercicio5() {
	theItCrowd := []string{
		"Maurice Moss",
		"Jen Barber",
		"Roy",
		"Richmond Avenal",
	}

	for i := 0; i < len(theItCrowd); i++ {
		fmt.Printf("%d.\t%s\n", i+1, theItCrowd[i])
	}
}
