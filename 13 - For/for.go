package main

import (
	"fmt"
	"time"
)

func main() {
	//for simples, passando apenas uma condição (pode representar um "while")
	var i int = 0
	for i < 10 {
		fmt.Println("Incrementando valor: ", i)
		i++
	}
	fmt.Println(i) //10

	//for casual, com inicialização, condição e incremento
	for j := 0; j < 5; j++ {
		fmt.Println("Incrementando valor de J: ", j)
	}

	fmt.Println("-----------------")

	//Para iterar em arrays, slices e maps, utilizaremos o "for range"
	//Range devolve o índice e o valor, ou a chave e o valor caso seja um map
	names := [3]string{"Gabriel", "Marcio", "Fabrício"}

	for index, name := range names {
		fmt.Println(index, name)
	}

	for _, name := range names {
		fmt.Println("Apenas um nome: ", name)
	}

	//Iterando strings, o "valor" retornado do range será o valor da tabela ASCII
	//Caso queira o valor real do caractere, utilize a função "string"
	for index, value := range "palavra" {
		fmt.Println(index, string(value))
	}

	//utilizando for em um map
	usuario := map[string]string{
		"nome":      "Gabriel",
		"sobrenome": "Gambôa",
	}

	for key, value := range usuario {
		fmt.Println(key, value)
	}

	//extra: while (true):
	for {
		fmt.Println("Cuidado pra não estourar a memória do teu PC!!!")
		time.Sleep(time.Second)
	}
}
