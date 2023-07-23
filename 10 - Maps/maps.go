package main

import "fmt"

func main() {
	//MAP: estrutura chave valor (map[tipoChave]tipoValor)
	usuario1 := map[string]string{
		"nome":      "Gabriel",
		"sobrenome": "Gamboa",
	}

	fmt.Println(usuario1)
	fmt.Println(usuario1["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro":  "Joana",
			"sobrenome": "Darc",
		},
		"curso": {
			"nome":   "ADS",
			"campus": "HTO",
		},
	}

	fmt.Println(usuario2)

	delete(usuario2, "curso") //deletar uma chave
	fmt.Println(usuario2)

	usuario2["curso"] = map[string]string{ //adicionar uma chave e um valor
		"nome":   "ADS",
		"campus": "HTO",
	}

	fmt.Println(usuario2)
}
