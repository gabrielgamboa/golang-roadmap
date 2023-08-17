package main

import (
	"fmt"
	"time"
)

func main() {
	var channel chan string
	channel = make(chan string)

	go escrever("Olá mundo!", channel)

	for {
		messageFromChannel, isOpen := <-channel

		if !isOpen {
			break
		}

		fmt.Println(messageFromChannel)
	}
}

func escrever(txt string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- txt
		time.Sleep(time.Second)
	}

	close(channel)
}
