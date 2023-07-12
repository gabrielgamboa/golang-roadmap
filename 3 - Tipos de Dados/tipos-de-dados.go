package main

import (
	"errors"
	"fmt"
)

func main() {
	numeroInteiro := 100000

	fmt.Println(numeroInteiro)

	var unsygnedInt uint = 1999 //não permite valor negativo

	fmt.Println(unsygnedInt)

	//ALIAS
	//rune = int32
	var numero3 rune = 1223456
	fmt.Println(numero3)

	//byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	var numeroReal2 float64 = 1234.56

	fmt.Println(numeroReal1, numeroReal2)

	var string1 string = "texto 1"
	string2 := "texto 2"
	fmt.Println(string1, string2)

	char := 'A' //vai devolver o valor inteiro da tabela ascii
	fmt.Println(char)

	//Valor 0 = valores iniciais das variáveis quando não são inicializadas

	var booleano bool //seu valor 0 é false
	fmt.Println(booleano)

	var erro1 error //seu valor 0 é <nil>
	var erro2 error = errors.New("Erro ao gerar variável")
	fmt.Println(erro1, erro2)
}
