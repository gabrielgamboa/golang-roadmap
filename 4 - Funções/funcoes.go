package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8, int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	multiplicacao := n1 * n2
	divisao := n1 / n2

	return soma, subtracao, multiplicacao, divisao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	f := func(txt string) {
		fmt.Println(txt)
	}

	f("Função declarada na variável")

	resultadoSoma, _, resultadoMultiplicacao, resultadoDivisao := calculosMatematicos(20, 5) //se eu não quiser usar determinado valor do conjunto de retornos, só utilizar um "_"
	fmt.Println(resultadoSoma, resultadoMultiplicacao, resultadoDivisao)
}
