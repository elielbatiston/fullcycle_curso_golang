package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

// Thread 1 é o próprio main
// Thread 2 é o garbage collection
// As minhas threads são as que eu coloco a palavra reservada go na frente da função.
func main() {
	// task("A")
	// task("B")

	//Thread 3
	go task("A")

	//Thread 4
	go task("B")

	//Thread 5
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(1 * time.Second)
		}
	}() //funcao anonima

	//Nada aqui então o programa vai sair. Nesse momento, se executarmos o programa, não veremos nada acontecendo.
	//Porem ao adicionarmos algo para segurar a execução iremos perceber as threads sendo executadas.
	time.Sleep(15 * time.Second)
}
