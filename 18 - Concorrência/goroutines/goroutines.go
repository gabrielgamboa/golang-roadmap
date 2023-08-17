package main

import "time"

func main() {
	// Concorrência != paralelismo

	go escrever("Olá mundo") //goroutine => executa e não espera terminar pra seguir o programa
	escrever("Programando em Go")
}

func escrever(txt string) {
	for {
		println(txt)
		time.Sleep(time.Second)
	}
}
