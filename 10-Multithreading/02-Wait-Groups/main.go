package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
		wg.Done() //subtrai 1 crédito daqui
	}
}

// Thread 1 é o próprio main
// Thread 2 é o garbage collection
// As minhas threads são as que eu coloco a palavra reservada go na frente da função.
func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(25) //adicionei 25 créditos aqui

	//Thread 3
	go task("A", &waitGroup)

	//Thread 4
	go task("B", &waitGroup)

	//Thread 5
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(1 * time.Second)
			waitGroup.Done() //subtrai 1 crédito daqui
		}
	}() //funcao anonima

	waitGroup.Wait() //Como sei que serão 25 iterações que terei, então eu iniciei um waitGroup, adicionei 25 iterações e agora eu peço pra ele esperar.
}
