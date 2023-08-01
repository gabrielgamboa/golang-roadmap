package main

import "fmt"

func sequencia() func() int {
	// var i int = 0

	// var incrementar func() int = func() int {
	// 	i += 1
	// 	return i
	// }
	i := 0

	incrementar := func() int {
		i += 1
		return i
	}

	return incrementar
}

func main() {
	fmt.Println("Closures!")

	//recebemos a função que referencia um valor que está contido dentro da função "sequencia"
	//seu estado será lembrado durante a execução do programa para cada chamada dessa função
	numero1 := sequencia()

	for i := 0; i < 5; i++ {
		fmt.Println(numero1())
	}

	fmt.Println("----------")

	numero2 := sequencia()

	for i := 0; i < 5; i++ {
		fmt.Println(numero2())
	}

}
