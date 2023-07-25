package main

import "fmt"

func calculosMatematicosProblematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return // mesma coisa que return soma, subtracao
}

func main() {
	soma, sub := calculosMatematicosProblematicos(10, 20)
	fmt.Println(soma, sub)
}
