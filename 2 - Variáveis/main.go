package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1"
	variavel2 := "Variável 2" //inferencia de tipo

	var (
		variavel3 string = "lalala"
		variavel4 string = "lalala"
	)

	variavel5, variavel6 := "Variável 5", "Variável 6"

	fmt.Println(variavel1)
	fmt.Println(variavel2)
	fmt.Println(variavel3, variavel4)
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante1" //não pode ser declarada com a sintaxe ":="
	fmt.Println(constante1)

}
