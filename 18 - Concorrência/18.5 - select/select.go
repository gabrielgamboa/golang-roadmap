package main

import (
	"fmt"
	"time"
)

func main() {
	channel1, channel2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			channel1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel2 <- "Canal 2"
		}
	}()

	for {
		// JEITO MAL OTIMIZADO, NÃO FUNCIONA DIREITO POIS ELE ESPERA O CANAL 2 RECEBER UMA MENSAGEM E REFAZER O LOOP
		// 	message1 := <-channel1
		// 	fmt.Println(message1)

		// 	message2 := <-channel2
		// 	fmt.Println(message2)

		//JEITO CORRETO UTILIZANDO O SELECT
		select {
		case message1 := <-channel1:
			fmt.Println(message1)
		case message2 := <-channel2:
			fmt.Println(message2)
		}
	}

}
