package main

import "fmt"

func main() {
	//função anonima = sem nome, executada diretamente

	//Sem retorno
	func(texto string) {
		fmt.Println(texto)
	}("Texto para imprimir")

	//Com retorno
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto) //Sprintf => formata o print e retorna o valor total em uma string
	}("Texto para imprimir")

	fmt.Println(retorno)
}
