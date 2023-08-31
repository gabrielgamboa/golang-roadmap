package main

import "fmt"

func main() {
	const MAX_SIZE_CHANNEL int = 45

	works := make(chan int, MAX_SIZE_CHANNEL)
	results := make(chan int, MAX_SIZE_CHANNEL)

	go worker(works, results) //se repetir a goroutine, ele processa mais rápido pois vão processar a mesma instância do canal

	for i := 0; i < MAX_SIZE_CHANNEL; i++ {
		works <- i
	}

	close(works)

	for i := 0; i < MAX_SIZE_CHANNEL; i++ {
		result := <-results
		fmt.Println(result)
	}

	close(results)
}

func worker(works <-chan int, results chan<- int) {
	for number := range works {
		results <- fibonacci(number)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
