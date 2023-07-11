package auxiliar

import "fmt"

func EscreverPublico() {
	fmt.Println("Escrevendo do arquivo auxiliar com método público.")
	escreverPrivado()
}
