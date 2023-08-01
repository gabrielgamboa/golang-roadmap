package main

import "fmt"

func alunoTaAprovado(n1, n2 int) bool {
	fmt.Println("Calculando média...")
	defer fmt.Println("Média calculada!")

	media := n1 + n2

	if media < 6 {
		return false
	}

	return true
}

func main() {
	// DEFER = adiar
	//Adia a execução de algo para o fim da função. Caso ela tiver return, ela será executada antes do return

	defer fmt.Println("Será executado por ultimo por conta do DEFER!!!")
	fmt.Println("Entendendo sobre a função defer")

	fmt.Println(alunoTaAprovado(5, 6))

	/*
		Entendendo sobre a função defer
		Calculando média...
		Média calculada!
		true
		Será executado por ultimo por conta do DEFER!!!
	*/
}
