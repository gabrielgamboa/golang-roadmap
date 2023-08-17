package main

import (
	"fmt"
	"time"
)

func main() {
	var channel chan string
	channel = make(chan string)

	go escrever("Olá mundo!", channel)

	//usando o canal e verificando se ele está aberto
	for {
		messageFromChannel, isOpen := <-channel //a variável só receberá o valor do canal quando tiver algo nele

		if !isOpen {
			break
		}

		fmt.Println(messageFromChannel)
	}

	// //usando sintaxe reduzida
	// for messageFromChannel := range channel {
	// 	fmt.Println(messageFromChannel)
	// }
}

func escrever(txt string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- txt
		time.Sleep(time.Second)
	}

	close(channel)
}
