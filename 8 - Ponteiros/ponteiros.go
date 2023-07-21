package main

import "fmt"

func main() {
	fmt.Println("ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	variavel1++
	fmt.Println(variavel1, variavel2) //atribuição por cópia, atualiza apenas a variável 1

	//PONTEIRO É UMA REFERENCIA DE MEMÓRIA

	var variavel3 int
	var ponteiro *int

	fmt.Println(variavel3, ponteiro) // 0 <nil>

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro, *ponteiro) //100, <enderecoMemoria> 100

	variavel3 = 150
	fmt.Println(variavel3, ponteiro, *ponteiro) //150, <enderecoMemoria> 150
}
