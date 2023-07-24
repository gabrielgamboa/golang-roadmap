package main

import (
	"fmt"
	"reflect"
)

func main() {
	//ARRAYS: tipos e tamanhos fixos
	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	//inferencia de tipo
	array2 := [4]string{"Posição 1, Posição 2, Posição 3, Posição 4"}
	fmt.Println(array2)

	//fixar tamanho do array com a quantidade de itens em sua inicialização
	array3 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(array3)

	//SLICES: tamanhos variáveis
	slice := []int{10, 15, 20}
	fmt.Println(slice)

	//IMPORTANTE: arrays e slices são tipos diferentes.
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18)
	fmt.Println(slice)

	//slice funciona como um ponteiro para essas posições
	slice2 := array2[1:3]
	fmt.Println(slice2) //[Posição 2, Posição 3]

	array2[1] = "Posição alterada"
	fmt.Println(slice2) //[Posição alterada, Posição 3]

	fmt.Println("-----------------------------")

	//Cria um ponteiro que aponta para um array interno no GO, de tamanho 4 e de length 3
	arrayComMake := make([]int, 3, 4)
	fmt.Println(arrayComMake)

	fmt.Println(len(arrayComMake)) //3
	fmt.Println(cap(arrayComMake)) //4

	arrayComMake = append(arrayComMake, 12)
	fmt.Println(arrayComMake)
	fmt.Println(len(arrayComMake)) //3
	fmt.Println(cap(arrayComMake)) //4

	arrayComMake = append(arrayComMake, 12) //Quando ultrapassar a capacidade máxima que era 4, ele dobra o valor da capacidade desse array interno
	fmt.Println(arrayComMake)
	fmt.Println(len(arrayComMake)) //3
	fmt.Println(cap(arrayComMake)) //4

}
