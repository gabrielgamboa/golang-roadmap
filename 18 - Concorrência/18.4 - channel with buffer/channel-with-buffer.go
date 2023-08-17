package main

import "fmt"

func main() {
	channel := make(chan string, 2)

	channel <- "Olá mundo"
	channel <- "Olá mundo 2"

	message := <-channel
	message2 := <-channel
	// message3 := <-channel ia dar erro, pois o tamanho do canal é de apenas duas strings

	fmt.Println(message)
	fmt.Println(message2)
}
