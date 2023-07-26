package main

import "fmt"

// "numeros" dentro da função vira um slice
// só pode ter um parâmetro variático por função, e tem que ser o último parâmetro caso tenha mais de um.
func somarNumeros(numeros ...int) int {
	var total int
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escreverTextoENumeros(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	soma := somarNumeros(1, 2, 3, 4, 5, 10)
	fmt.Println(soma)

	escreverTextoENumeros("Texto legal! ", 10, 12, 14)
}
