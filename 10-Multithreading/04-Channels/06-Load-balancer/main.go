package main

import (
	"fmt"
	"time"
)

// Thread 1 é o próprio main
// Thread 2 é o garbage collection
func main() {
	data := make(chan int)
	qtdWorkers := 100000

	for i := 0; i < qtdWorkers; i++ {
		go worker(i, data)
	}

	for i := 0; i < 1000000; i++ { //envia informações para o canal
		data <- i
	}
}

func worker(workerId int, data chan int) {
	for x := range data { //Le as informações do canal
		fmt.Printf("Worker %d receive %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}
