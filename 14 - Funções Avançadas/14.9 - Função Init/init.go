package main

import "fmt"

var number int

func main() {
	//inicializou na init
	fmt.Println(number)
}

// pode ter uma por arquivo, diferente da main que possui uma por pacote
func init() {
	number = 10
}
