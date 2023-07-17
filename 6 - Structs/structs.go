package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo de structs!")

	var user usuario

	user.idade = 12
	user.nome = "Gabriel"

	fmt.Println(user) // { Gabriel 12 }

	//passar os valores na ordem correta (todos os campos serão obrigatórios)
	enderecoDoUser2 := endereco{"Rua dos bobos", 0}
	user2 := usuario{"Gustavo", 20, enderecoDoUser2} //inferencia de structs no GO
	fmt.Println(user2)

	//passar apenas alguns valores (usando chave e valor)
	user3 := usuario{idade: 22}
	fmt.Println(user3)

}
