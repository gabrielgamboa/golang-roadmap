package main

import (
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	goroutines_in_program := 2

	waitGroup.Add(goroutines_in_program)

	go func() {
		escrever("Olá mundo")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("Programando em Go e utilizando waitgroup para lidar com sincronismo de goroutines")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait() // 0

}

func escrever(txt string) {
	for i := 0; i < 5; i++ {
		println(txt)
		time.Sleep(time.Second)
	}
}
