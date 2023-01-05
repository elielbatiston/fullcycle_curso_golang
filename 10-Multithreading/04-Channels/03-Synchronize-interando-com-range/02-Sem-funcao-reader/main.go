package main

import "fmt"

// Thread 1 é o próprio main
// Thread 2 é o garbage collection
func main() {
	ch := make(chan int)
	go publish(ch)

	for x := range ch {
		fmt.Printf("Received %d\n", x)
		//como estamos num loop lendo o canal, assim que finalizar as 10 iteracoes, vai dar deadlock na aplicação.
		//para resolver isso, no publish sinalizamos para o go fechar o canal.
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) //fechando o canal para evitarmos deadlock
}
