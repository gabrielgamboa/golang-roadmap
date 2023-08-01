package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!!!")
	}
}

func alunoTaAprovado(n1, n2 int) bool {
	defer recuperarExecucao()

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	//Panic mata a execução do programa e chama as funções que estão marcadas com DEFER
	panic("A MEDIA É 6!!!!")
}

func main() {
	fmt.Println(alunoTaAprovado(6, 6))
	fmt.Println("Pós execução")
}
