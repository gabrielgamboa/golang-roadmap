package main

import (
	"fmt"
	"module/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.EscreverPublico()

	//utilizando biblioteca externa
	error := checkmail.ValidateFormat("gabriel")

	if error != nil {
		fmt.Println(error)
	}
}
