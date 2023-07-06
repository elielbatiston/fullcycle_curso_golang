package main

import "fmt"

// Thread 1 é o próprio main
// Thread 2 é o garbage collection
func main() {
	canal := make(chan string) //Canal vazio

	//Thread 3
	go func() {
		canal <- "Olá Mundo!" //Canal cheio
	}()

	//Thread 1
	msg := <-canal //Canal esvazia
	fmt.Println(msg)
}
