package main

import "fmt"

func reverseNumberSign(number int) int {
	return number * -1
}

func reverseNumberSignWithPointer(number *int) int {
	return *number * -1
}

func main() {
	number1 := 10
	newNumber1 := reverseNumberSign(number1)
	fmt.Println(newNumber1)
	fmt.Println(number1)

	number2 := 20
	fmt.Println(reverseNumberSignWithPointer(&number2)) //passa a referencia de memória para a função
	fmt.Println(number2)

}
