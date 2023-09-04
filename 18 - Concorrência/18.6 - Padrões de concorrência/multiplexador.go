package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := multiplexador(write("Hello world!!!"), write("Hello go!!!!!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func multiplexador(channel1, channel2 <-chan string) <-chan string {
	channelMultiplex := make(chan string)

	go func() {
		for {
			select {
			case message := <-channel1:
				channelMultiplex <- message
			case message := <-channel2:
				channelMultiplex <- message
			}
		}
	}()

	return channelMultiplex
}

func write(txt string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Message received: %s", txt)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return channel
}
