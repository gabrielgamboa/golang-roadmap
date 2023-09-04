package main

import (
	"fmt"
	"time"
)

func main() {
	channel := write("Hello Go!")

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

// função que encapsula gorountines e criação de canais.
// retorna um canal que recebe dados e são lidos na funcão main
func write(txt string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Message received: %s", txt)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return channel
}
