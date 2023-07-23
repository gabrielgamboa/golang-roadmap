package main

import "fmt"

func main() {
	numero := 6

	if numero%2 == 0 {
		fmt.Println("Numero par")
	} else {
		fmt.Println("Numero impar")
	}

	if outroNumero := numero; outroNumero > 0 { //outroNumero existe apenas no contexto do if
		fmt.Println("Numero é maior que zero")
	} else if outroNumero < -10 {
		fmt.Println("Numero menor que -10")
	} else {
		fmt.Println("Numero é menor que zero e maior que -10")
	}

	// fmt.Println(outroNumero) NÃO FUNCIONA
}
