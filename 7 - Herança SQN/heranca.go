package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	pessoa := pessoa{"Gabriel", "Gambôa", 22, 80}

	fmt.Println(pessoa)

	estudante := estudante{pessoa, "ADS", "IFSP HTO"}
	fmt.Println(estudante)
}
